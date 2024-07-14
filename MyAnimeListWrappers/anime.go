package MyAnimeListWrappers

import (
	"encoding/json"
	"fmt"
	"net/http"

	MALModels "github.com/HikaruHokkyokusei/MAL-SDK/MyAnimeListModels"
)

var (
	nilAnime = MALModels.Anime{}
)

func (c MyAnimeListClient) GetAnime(animeId int64) (MALModels.Anime, error) {
	resp, err := c.request(
		http.MethodGet, fmt.Sprintf("/anime/%d", animeId),
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
		http.MethodGet, "/anime",
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
		return convertIntoEntityList(animeList), nil
	} else {
		return nil, err
	}
}

func (c MyAnimeListClient) GetAnimeRanking(rankingCategory AnimeRankingCategory, limit int64, offset int64) ([]MALModels.Anime, error) {
	if limit < MinResultListSize || limit > MaxResultListSize {
		limit = DefaultResultListSize
	}

	resp, err := c.request(
		http.MethodGet, "/anime/ranking",
		map[string]string{
			"ranking_type": rankingCategory,
			"limit":        fmt.Sprintf("%d", limit),
			"offset":       fmt.Sprintf("%d", offset),
			"fields":       MALModels.AllAnimeFields(),
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
		return convertIntoEntityList(animeList), nil
	} else {
		return nil, err
	}
}

func (c MyAnimeListClient) GetSeasonalAnime(year int64, season MalAnimeSeason, limit int64, offset int64) ([]MALModels.Anime, error) {
	if limit < MinResultListSize || limit > MaxResultListSize {
		limit = DefaultResultListSize
	}

	resp, err := c.request(
		http.MethodGet, fmt.Sprintf("/anime/season/%d/%s", year, season),
		map[string]string{
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
		return convertIntoEntityList(animeList), nil
	} else {
		return nil, err
	}
}

func (c MyAnimeListClient) GetSuggestedAnime(limit int64, offset int64) ([]MALModels.Anime, error) {
	if limit < MinResultListSize || limit > MaxResultListSize {
		limit = DefaultResultListSize
	}

	resp, err := c.request(
		http.MethodGet, "/anime/suggestions",
		map[string]string{
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
		return convertIntoEntityList(animeList), nil
	} else {
		return nil, err
	}
}
