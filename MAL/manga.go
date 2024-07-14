package MyAnimeListSDK

import (
	"encoding/json"
	"fmt"
	"net/http"

	MALModels "github.com/HikaruHokkyokusei/AnimeTree/MAL/Models"
)

var (
	nilManga = MALModels.Manga{}
)

func (c MyAnimeListClient) GetManga(mangaId int64) (MALModels.Manga, error) {
	resp, err := c.request(
		http.MethodGet, fmt.Sprintf("/manga/%d", mangaId),
		map[string]string{
			"fields": MALModels.AllMangaFields(),
		}, "",
	)
	if err != nil {
		return nilManga, err
	}

	if temp, err := json.Marshal(resp); err == nil {
		manga := MALModels.Manga{}
		return manga, json.Unmarshal(temp, &manga)
	} else {
		return nilManga, err
	}
}
