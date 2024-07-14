package MyAnimeListWrappers

import (
	"encoding/json"
	"fmt"
	"net/http"

	MALModels "github.com/HikaruHokkyokusei/MAL-SDK/MyAnimeListModels"
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

func (c MyAnimeListClient) SearchManga(mangaName string, limit int64, offset int64) ([]MALModels.Manga, error) {
	if limit < MinResultListSize || limit > MaxResultListSize {
		limit = DefaultResultListSize
	}

	resp, err := c.request(
		http.MethodGet, "/manga",
		map[string]string{
			"q":      mangaName,
			"limit":  fmt.Sprintf("%d", limit),
			"offset": fmt.Sprintf("%d", offset),
			"fields": MALModels.AllMangaFields(),
		}, "",
	)
	if err != nil {
		return nil, err
	}

	mangaList := MALModels.ResponsePage[MALModels.Manga]{}
	if temp, err := json.Marshal(resp); err == nil {
		if err := json.Unmarshal(temp, &mangaList); err != nil {
			return nil, err
		}
		return convertIntoEntityList(mangaList), nil
	} else {
		return nil, err
	}
}

func (c MyAnimeListClient) GetMangaRanking(rankingCategory MangaRankingCategory, limit int64, offset int64) ([]MALModels.Manga, error) {
	if limit < MinResultListSize || limit > MaxResultListSize {
		limit = DefaultResultListSize
	}

	resp, err := c.request(
		http.MethodGet, "/manga/ranking",
		map[string]string{
			"ranking_type": rankingCategory,
			"limit":        fmt.Sprintf("%d", limit),
			"offset":       fmt.Sprintf("%d", offset),
			"fields":       MALModels.AllMangaFields(),
		}, "",
	)
	if err != nil {
		return nil, err
	}

	mangaList := MALModels.ResponsePage[MALModels.Manga]{}
	if temp, err := json.Marshal(resp); err == nil {
		if err := json.Unmarshal(temp, &mangaList); err != nil {
			return nil, err
		}
		return convertIntoEntityList(mangaList), nil
	} else {
		return nil, err
	}
}
