package MALModels

type AlternativeTitles struct {
	En       *string   `json:"en"`
	Ja       *string   `json:"ja"`
	Synonyms *[]string `json:"synonyms"`
}

type Broadcast struct {
	DayOfTheWeek *string `json:"day_of_the_week"`
	StartTime    *string `json:"start_time"`
}

type Genre struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name"`
}

type Picture struct {
	Large  *string `json:"large"`
	Medium *string `json:"medium"`
}

type AnimeNode struct {
	Node               *Anime `json:"node"`
	NumRecommendations *int64 `json:"num_recommendations"`
}

type StartSession struct {
	Season *string `json:"season"`
	Year   *int64  `json:"year"`
}

type Studio struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name"`
}

type Anime struct {
	AlternativeTitles      *AlternativeTitles `json:"alternative_titles"`
	AverageEpisodeDuration *int64             `json:"average_episode_duration"`
	Broadcast              *Broadcast         `json:"broadcast"`
	CreatedAt              *string            `json:"created_at"`
	EndDate                *string            `json:"end_date"`
	Genres                 *[]Genre           `json:"genres"`
	Id                     *int64             `json:"id"`
	MainPicture            *Picture           `json:"main_picture"`
	Mean                   *float64           `json:"mean"`
	MediaType              *string            `json:"media_type"`
	NSFW                   *string            `json:"nsfw"`
	NumEpisodes            *int64             `json:"num_episodes"`
	NumFavorites           *int64             `json:"num_favorites"`
	NumListUsers           *int64             `json:"num_list_users"`
	NumScoringUsers        *int64             `json:"num_scoring_users"`
	Pictures               *[]Picture         `json:"pictures"`
	Popularity             *int64             `json:"popularity"`
	Rank                   *int64             `json:"rank"`
	Rating                 *string            `json:"rating"`
	Recommendations        *[]AnimeNode       `json:"recommendations"`
	RelatedAnime           *[]AnimeNode       `json:"related_anime"`
	Source                 *string            `json:"source"`
	StartDate              *string            `json:"start_date"`
	StartSession           *StartSession      `json:"start_session"`
	Status                 *string            `json:"status"`
	Studios                *[]Studio          `json:"studios"`
	Synopsis               *string            `json:"synopsis"`
	Title                  *string            `json:"title"`
	UpdatedAt              *string            `json:"updated_at"`
}
