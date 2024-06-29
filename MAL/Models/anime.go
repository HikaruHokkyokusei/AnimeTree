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

type MyListStatus struct {
	FinishDate         *string `json:"finish_date"`
	IsRewatching       *bool   `json:"is_rewatching"`
	NumEpisodesWatched *int64  `json:"num_episodes_watched"`
	Score              *int64  `json:"score"`
	StartDate          *string `json:"start_date"`
	Status             *string `json:"status"`
	UpdatedAt          *string `json:"updated_at"`
}

type RecommendationNode struct {
	Node               *AnimeNode `json:"node"`
	NumRecommendations *int64     `json:"num_recommendations"`
}

type RelatedAnimeNode struct {
	Node                  *AnimeNode `json:"node"`
	RelationType          *string    `json:"relation_type"`
	RelationTypeFormatted *string    `json:"relation_type_formatted"`
}

type StartSession struct {
	Season *string `json:"season"`
	Year   *int64  `json:"year"`
}

type Studio struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name"`
}

type AnimeNode struct {
	AlternativeTitles      *AlternativeTitles `json:"alternative_titles"`
	AverageEpisodeDuration *int64             `json:"average_episode_duration"`
	CreatedAt              *string            `json:"created_at"`
	EndDate                *string            `json:"end_date"`
	Genres                 *[]Genre           `json:"genres"`
	Id                     *int64             `json:"id"`
	MainPicture            *Picture           `json:"main_picture"`
	Mean                   *float64           `json:"mean"`
	MediaType              *string            `json:"media_type"`
	MyListStatus           *MyListStatus      `json:"my_list_status"`
	StartDate              *string            `json:"start_date"`
	Title                  *string            `json:"title"`
}

type Anime struct {
	AlternativeTitles      *AlternativeTitles    `json:"alternative_titles"`
	AverageEpisodeDuration *int64                `json:"average_episode_duration"`
	Broadcast              *Broadcast            `json:"broadcast"`
	CreatedAt              *string               `json:"created_at"`
	EndDate                *string               `json:"end_date"`
	Genres                 *[]Genre              `json:"genres"`
	Id                     *int64                `json:"id"`
	MainPicture            *Picture              `json:"main_picture"`
	Mean                   *float64              `json:"mean"`
	MediaType              *string               `json:"media_type"`
	MyListStatus           *MyListStatus         `json:"my_list_status"`
	NSFW                   *string               `json:"nsfw"`
	NumEpisodes            *int64                `json:"num_episodes"`
	NumFavorites           *int64                `json:"num_favorites"`
	NumListUsers           *int64                `json:"num_list_users"`
	NumScoringUsers        *int64                `json:"num_scoring_users"`
	Pictures               *[]Picture            `json:"pictures"`
	Popularity             *int64                `json:"popularity"`
	Rank                   *int64                `json:"rank"`
	Rating                 *string               `json:"rating"`
	Recommendations        *[]RecommendationNode `json:"recommendations"`
	RelatedAnime           *[]RelatedAnimeNode   `json:"related_anime"`
	Source                 *string               `json:"source"`
	StartDate              *string               `json:"start_date"`
	StartSession           *StartSession         `json:"start_session"`
	Status                 *string               `json:"status"`
	Studios                *[]Studio             `json:"studios"`
	Synopsis               *string               `json:"synopsis"`
	Title                  *string               `json:"title"`
	UpdatedAt              *string               `json:"updated_at"`
}
