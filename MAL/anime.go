package MyAnimeListSDK

import (
	"encoding/json"
	"fmt"

	MALModels "github.com/HikaruHokkyokusei/AnimeTree/MAL/Models"
)

var (
	nilAnime = MALModels.Anime{}
)

func getAnimeList(animeDataPage MALModels.ResponsePage[MALModels.Anime]) []MALModels.Anime {
	var animeList []MALModels.Anime
	for _, node := range *animeDataPage.Data {
		animeList = append(animeList, node.Node)
	}
	return animeList
}

func (c MyAnimeListClient) GetAnime(animeId int64) (MALModels.Anime, error) {
	resp, err := c.request(
		GET, fmt.Sprintf("/anime/%d", animeId),
		map[string]string{
			"fields": MALModels.AllAnimeFields(),
		}, "",
	)
	if err != nil {
		return nilAnime, err
	}

	if temp, err := json.Marshal(resp); err == nil {
		anime := MALModels.Anime{}
		return anime, json.Unmarshal(temp, &anime)
	} else {
		return nilAnime, err
	}
}

func (c MyAnimeListClient) SearchAnime(animeName string, limit int64, offset int64) ([]MALModels.Anime, error) {
	if limit < MinResultListSize || limit > MaxResultListSize {
		limit = DefaultResultListSize
	}

	resp, err := c.request(
		GET, "/anime",
		map[string]string{
			"q":      animeName,
			"limit":  fmt.Sprintf("%d", limit),
			"offset": fmt.Sprintf("%d", offset),
			"fields": MALModels.AllAnimeFields(),
		}, "",
	)
	if err != nil {
		return nil, err
	}

	animeList := MALModels.ResponsePage[MALModels.Anime]{}
	if temp, err := json.Marshal(resp); err == nil {
		if err := json.Unmarshal(temp, &animeList); err != nil {
			return nil, err
		}
		return getAnimeList(animeList), nil
	} else {
		return nil, err
	}
}
