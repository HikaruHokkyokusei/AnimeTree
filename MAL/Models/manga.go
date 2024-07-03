package MALModels

type Author struct {
	Node *struct {
		Id        *int64  `json:"id,omitempty"`
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
	} `json:"node,omitempty"`
	Role *string `json:"role,omitempty"`
}

type MangaMyListStatus struct {
	FinishDate      *string `json:"finish_date,omitempty"`
	IsRereading     *bool   `json:"is_rereading,omitempty"`
	NumChaptersRead *int64  `json:"num_chapters_read,omitempty"`
	NumVolumesRead  *int64  `json:"num_volumes_read,omitempty"`
	Score           *int64  `json:"score,omitempty"`
	StartDate       *string `json:"start_date,omitempty"`
	Status          *string `json:"status,omitempty"`
	UpdatedAt       *string `json:"updated_at,omitempty"`
}

type Serialization struct {
	Node *IdNameEntity `json:"node,omitempty"`
}

type SubMangaNode struct {
	AlternativeTitles *AlternativeTitles `json:"alternative_titles,omitempty"`
	Authors           *[]Author          `json:"authors,omitempty"`
	CreatedAt         *string            `json:"created_at,omitempty"`
	EndDate           *string            `json:"end_date,omitempty"`
	Genres            *[]IdNameEntity    `json:"genres,omitempty"`
	Id                *int64             `json:"id,omitempty"`
	MainPicture       *Picture           `json:"main_picture,omitempty"`
	Mean              *float64           `json:"mean,omitempty"`
	MediaType         *string            `json:"media_type,omitempty"`
	MyListStatus      *MangaMyListStatus `json:"my_list_status,omitempty"`
	NSFW              *string            `json:"nsfw,omitempty"`
	NumChapters       *int64             `json:"num_chapters,omitempty"`
	NumFavorites      *int64             `json:"num_favorites,omitempty"`
	NumListUsers      *int64             `json:"num_list_users,omitempty"`
	NumScoringUsers   *int64             `json:"num_scoring_users,omitempty"`
	NumVolumes        *int64             `json:"num_volumes,omitempty"`
	Popularity        *int64             `json:"popularity,omitempty"`
	Rank              *int64             `json:"rank,omitempty"`
	StartDate         *string            `json:"start_date,omitempty"`
	Status            *string            `json:"status,omitempty"`
	Synopsis          *string            `json:"synopsis,omitempty"`
	Title             *string            `json:"title,omitempty"`
	UpdatedAt         *string            `json:"updated_at,omitempty"`
}

type Manga struct {
	AlternativeTitles *AlternativeTitles                  `json:"alternative_titles,omitempty"`
	Authors           *[]Author                           `json:"authors,omitempty"`
	Background        *string                             `json:"background,omitempty"`
	CreatedAt         *string                             `json:"created_at,omitempty"`
	EndDate           *string                             `json:"end_date,omitempty"`
	Genres            *[]IdNameEntity                     `json:"genres,omitempty"`
	Id                *int64                              `json:"id,omitempty"`
	MainPicture       *Picture                            `json:"main_picture,omitempty"`
	Mean              *float64                            `json:"mean,omitempty"`
	MediaType         *string                             `json:"media_type,omitempty"`
	MyListStatus      *MangaMyListStatus                  `json:"my_list_status,omitempty"`
	NSFW              *string                             `json:"nsfw,omitempty"`
	NumChapters       *int64                              `json:"num_chapters,omitempty"`
	NumFavorites      *int64                              `json:"num_favorites,omitempty"`
	NumListUsers      *int64                              `json:"num_list_users,omitempty"`
	NumScoringUsers   *int64                              `json:"num_scoring_users,omitempty"`
	NumVolumes        *int64                              `json:"num_volumes,omitempty"`
	Pictures          *[]Picture                          `json:"pictures,omitempty"`
	Popularity        *int64                              `json:"popularity,omitempty"`
	Rank              *int64                              `json:"rank,omitempty"`
	Recommendations   *[]RecommendationNode[SubMangaNode] `json:"recommendations,omitempty"`
	RelatedAnime      *[]RelatedEntityNode[SubAnimeNode]  `json:"related_anime,omitempty"`
	RelatedManga      *[]RelatedEntityNode[SubMangaNode]  `json:"related_manga,omitempty"`
	Serialization     *[]Serialization                    `json:"serialization,omitempty"`
	StartDate         *string                             `json:"start_date,omitempty"`
	Status            *string                             `json:"status,omitempty"`
	Synopsis          *string                             `json:"synopsis,omitempty"`
	Title             *string                             `json:"title,omitempty"`
	UpdatedAt         *string                             `json:"updated_at,omitempty"`
}
