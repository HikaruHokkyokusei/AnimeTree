package MALModels

type Broadcast struct {
	DayOfTheWeek *string `json:"day_of_the_week,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
}

type AnimeMyListStatus struct {
	FinishDate         *string `json:"finish_date,omitempty"`
	IsRewatching       *bool   `json:"is_rewatching,omitempty"`
	NumEpisodesWatched *int64  `json:"num_episodes_watched,omitempty"`
	Score              *int64  `json:"score,omitempty"`
	StartDate          *string `json:"start_date,omitempty"`
	Status             *string `json:"status,omitempty"`
	UpdatedAt          *string `json:"updated_at,omitempty"`
}

type StartSeason struct {
	Season *string `json:"season,omitempty"`
	Year   *int64  `json:"year,omitempty"`
}

type AnimeStatistics struct {
	Status *struct {
		Watching    *string `json:"watching,omitempty"`
		Completed   *string `json:"completed,omitempty"`
		OnHold      *string `json:"on_hold,omitempty"`
		Dropped     *string `json:"dropped,omitempty"`
		PlanToWatch *string `json:"plan_to_watch,omitempty"`
	} `json:"status,omitempty"`
	NumListUsers *int64 `json:"num_list_users,omitempty"`
}

type SubAnimeNode struct {
	AlternativeTitles      *AlternativeTitles `json:"alternative_titles,omitempty"`
	AverageEpisodeDuration *int64             `json:"average_episode_duration,omitempty"`
	Broadcast              *Broadcast         `json:"broadcast,omitempty"`
	CreatedAt              *string            `json:"created_at,omitempty"`
	EndDate                *string            `json:"end_date,omitempty"`
	Genres                 *[]IdNameEntity    `json:"genres,omitempty"`
	Id                     *int64             `json:"id,omitempty"`
	MainPicture            *Picture           `json:"main_picture,omitempty"`
	Mean                   *float64           `json:"mean,omitempty"`
	MediaType              *string            `json:"media_type,omitempty"`
	MyListStatus           *AnimeMyListStatus `json:"my_list_status,omitempty"`
	NSFW                   *string            `json:"nsfw,omitempty"`
	NumEpisodes            *int64             `json:"num_episodes,omitempty"`
	NumFavorites           *int64             `json:"num_favorites,omitempty"`
	NumListUsers           *int64             `json:"num_list_users,omitempty"`
	NumScoringUsers        *int64             `json:"num_scoring_users,omitempty"`
	Popularity             *float64           `json:"popularity,omitempty"`
	Rank                   *int64             `json:"rank,omitempty"`
	Rating                 *string            `json:"rating,omitempty"`
	Source                 *string            `json:"source,omitempty"`
	StartDate              *string            `json:"start_date,omitempty"`
	StartSeason            *StartSeason       `json:"start_season,omitempty"`
	Status                 *string            `json:"status,omitempty"`
	Studios                *[]IdNameEntity    `json:"studios,omitempty"`
	Synopsis               *string            `json:"synopsis,omitempty"`
	Title                  *string            `json:"title,omitempty"`
	UpdatedAt              *string            `json:"updated_at,omitempty"`
}

type Anime struct {
	AlternativeTitles      *AlternativeTitles                  `json:"alternative_titles,omitempty"`
	AverageEpisodeDuration *int64                              `json:"average_episode_duration,omitempty"`
	Background             *string                             `json:"background,omitempty"`
	Broadcast              *Broadcast                          `json:"broadcast,omitempty"`
	CreatedAt              *string                             `json:"created_at,omitempty"`
	EndDate                *string                             `json:"end_date,omitempty"`
	Genres                 *[]IdNameEntity                     `json:"genres,omitempty"`
	Id                     *int64                              `json:"id,omitempty"`
	MainPicture            *Picture                            `json:"main_picture,omitempty"`
	Mean                   *float64                            `json:"mean,omitempty"`
	MediaType              *string                             `json:"media_type,omitempty"`
	MyListStatus           *AnimeMyListStatus                  `json:"my_list_status,omitempty"`
	NSFW                   *string                             `json:"nsfw,omitempty"`
	NumEpisodes            *int64                              `json:"num_episodes,omitempty"`
	NumFavorites           *int64                              `json:"num_favorites,omitempty"`
	NumListUsers           *int64                              `json:"num_list_users,omitempty"`
	NumScoringUsers        *int64                              `json:"num_scoring_users,omitempty"`
	Pictures               *[]Picture                          `json:"pictures,omitempty"`
	Popularity             *int64                              `json:"popularity,omitempty"`
	Rank                   *int64                              `json:"rank,omitempty"`
	Rating                 *string                             `json:"rating,omitempty"`
	Recommendations        *[]RecommendationNode[SubAnimeNode] `json:"recommendations,omitempty"`
	RelatedAnime           *[]RelatedEntityNode[SubAnimeNode]  `json:"related_anime,omitempty"`
	RelatedManga           *[]RelatedEntityNode[SubMangaNode]  `json:"related_manga,omitempty"`
	Source                 *string                             `json:"source,omitempty"`
	StartDate              *string                             `json:"start_date,omitempty"`
	StartSeason            *StartSeason                        `json:"start_season,omitempty"`
	Statistics             *any                                `json:"statistics,omitempty"`
	Status                 *string                             `json:"status,omitempty"`
	Studios                *[]IdNameEntity                     `json:"studios,omitempty"`
	Synopsis               *string                             `json:"synopsis,omitempty"`
	Title                  *string                             `json:"title,omitempty"`
	UpdatedAt              *string                             `json:"updated_at,omitempty"`
}
