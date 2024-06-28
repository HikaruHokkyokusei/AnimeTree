package mal_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	Fiber "github.com/gofiber/fiber/v2"
	UUID "github.com/google/uuid"
	PKCE "github.com/nirasan/go-oauth-pkce-code-verifier"
	MAL "github.com/nstratos/go-myanimelist/mal"
	Browser "github.com/pkg/browser"
	OAuth2 "golang.org/x/oauth2"
)

var AuthFile *os.File
var AuthToken *OAuth2.Token

func generateAuthServerAndURL(authConfig OAuth2.Config, callbackPath string, codeQueryKey string, tokenChannel chan *OAuth2.Token) (*Fiber.App, string, error) {
	state, err := UUID.NewV7()
	if err != nil {
		return nil, "", fmt.Errorf("error generating state: %v", err)
	}

	pkceVerifierCode, err := PKCE.CreateCodeVerifier()
	if err != nil {
		return nil, "", err
	}
	pkceChallengeCode := pkceVerifierCode.CodeChallengePlain()

	server := Fiber.New()
	server.Get(callbackPath, func(ctx *Fiber.Ctx) error {
		code := ctx.Query(codeQueryKey)
		if code == "" {
			_ = server.Shutdown()
			tokenChannel <- nil
			panic("Auth code is empty")
		} else {
			fmt.Println("User Authorization Successful. Obtaining token...")
		}

		authToken, err := authConfig.Exchange(
			context.Background(), code,
			OAuth2.SetAuthURLParam("code_verifier", pkceVerifierCode.String()),
		)
		if err != nil {
			_ = ctx.SendStatus(http.StatusInternalServerError)
			_ = server.Shutdown()
			tokenChannel <- nil
			panic(err)
		}

		_ = ctx.SendStatus(http.StatusOK)
		tokenChannel <- authToken
		return nil
	})

	url := authConfig.AuthCodeURL(
		state.String(),
		OAuth2.SetAuthURLParam("code_challenge", pkceChallengeCode),
		OAuth2.SetAuthURLParam("code_challenge_method", "plain"),
	)
	return server, url, nil
}

func authenticateUserViaServer(authConfig OAuth2.Config, serverPortNumber string) (*OAuth2.Token, error) {
	tokenChannel := make(chan *OAuth2.Token)

	authServer, authURL, err := generateAuthServerAndURL(authConfig, "/callback", "code", tokenChannel)
	if err != nil {
		return nil, fmt.Errorf("failed to generate auth URL")
	}

	go func() {
		if err := authServer.Listen(":" + serverPortNumber); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := authServer.ShutdownWithTimeout(5 * time.Second); err != nil {
			fmt.Println("Error shutting down server")
		}
	}()

	fmt.Println("Waiting for user authorization...")
	if err := Browser.OpenURL(authURL); err != nil {
		return nil, err
	}

	token := <-tokenChannel
	fmt.Println("Successfully obtained token.")
	return token, nil
}

func BuildMALClient(clientId string, clientSecret string) *MAL.Client {
	MALAuthConfig := OAuth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint: OAuth2.Endpoint{
			AuthURL:   "https://myanimelist.net/v1/oauth2/authorize",
			AuthStyle: OAuth2.AuthStyleInParams,
			TokenURL:  "https://myanimelist.net/v1/oauth2/token",
		},
		Scopes:      []string{"read:users"},
		RedirectURL: "http://localhost:8080/callback",
	}

	if AuthToken == nil {
		authToken, err := authenticateUserViaServer(MALAuthConfig, "8080")
		fmt.Println("Authentication Successful")
		if err != nil {
			panic(err)
		}
		AuthToken = authToken
	}

	return MAL.NewClient(MALAuthConfig.Client(context.Background(), AuthToken))
}

func init() {
	var err error

	if AuthFile, err = os.OpenFile("./authToken.store", os.O_CREATE|os.O_RDWR, 0755); err != nil {
		panic("Unable to access auth store file")
	}

	buffer, err := io.ReadAll(AuthFile)
	if len(buffer) > 0 && json.Valid(buffer) {
		AuthToken = &OAuth2.Token{}
		err := json.Unmarshal(buffer, AuthToken)
		if err != nil {
			panic(err)
		}
	}
}

func Exit() {
	if AuthToken != nil {
		authToken, err := json.MarshalIndent(AuthToken, "", "  ")

		if err == nil {
			_ = AuthFile.Truncate(0)
			_, _ = AuthFile.Seek(0, 0)
			_, _ = AuthFile.Write(authToken)
			_, _ = AuthFile.Write([]byte("\n"))
		}
	}

	_ = AuthFile.Close()
}
