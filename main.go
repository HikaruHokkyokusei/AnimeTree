package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HikaruHokkyokusei/AnimeTree/MAL"
)

var (
	malClientSecret string
	malClientId     string
)

func logic() {
	malClient := MyAnimeListSDK.BuildClient(malClientId, malClientSecret)
	fmt.Println(malClient)
}

func init() {
	malClientId = os.Getenv("MAL_CLIENT_ID")
	malClientSecret = os.Getenv("MAL_CLIENT_SECRET")

	if malClientId == "" || malClientSecret == "" {
		panic("ClientId and ClientSecret env variables not set")
	}
}

func main() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			signalChannel <- syscall.SIGINT
		}()
		logic()
	}()

	<-signalChannel
	shutdownHandler()
}

func shutdownHandler() {
	MyAnimeListSDK.Exit()
	os.Exit(0)
}
