package MyAnimeListSDK

import (
	"encoding/json"
	"fmt"
	"strings"

	MALModels "github.com/HikaruHokkyokusei/AnimeTree/MAL/Models"
)

var (
	animeFieldsTier1 = []string{
		"alternative_titles",
		"average_episode_duration",
		"created_at",
		"end_date",
		"genres",
		"id",
		"main_picture",
		"mean",
		"media_type",
		"my_list_status{finish_date,is_rewatching,num_episodes_watched,score,start_date,status,updated_at}",
		"start_date",
		"title",
	}
	animeFieldsTier2 = []string{
		"background",
		"broadcast",
		"nsfw",
		"num_episodes",
		"num_favorites",
		"num_list_users",
		"num_scoring_users",
		"pictures",
		"popularity",
		"rank",
		"rating",
		"source",
		"start_season",
		"status",
		"studios",
		"synopsis",
		"updated_at",
	}
	animeFieldsTier3 = []string{
		"recommendations{%s}",
		"related_anime{%s}",
	}
	nilAnime = MALModels.Anime{}
)

func getAllFieldsForAnime() string {
	t1 := strings.Join(animeFieldsTier1, ",")
	t2 := strings.Join(animeFieldsTier2, ",")
	t3 := fmt.Sprintf(strings.Join(animeFieldsTier3, ","), t1, t1)
	return strings.Join([]string{t1, t2, t3}, ",")
}

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
			"fields": getAllFieldsForAnime(),
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
			"fields": getAllFieldsForAnime(),
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
