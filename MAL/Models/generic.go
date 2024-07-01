package MALModels

type AlternativeTitles struct {
	En       *string   `json:"en,omitempty"`
	Ja       *string   `json:"ja,omitempty"`
	Synonyms *[]string `json:"synonyms,omitempty"`
}

type IdNameEntity struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Picture struct {
	Large  *string `json:"large,omitempty"`
	Medium *string `json:"medium,omitempty"`
}

type RelatedEntityNode[SubEntityNode SubAnimeNode | SubMangaNode] struct {
	Node                  *SubEntityNode `json:"node,omitempty"`
	RelationType          *string        `json:"relation_type,omitempty"`
	RelationTypeFormatted *string        `json:"relation_type_formatted,omitempty"`
}

type RecommendationNode[SubEntityNode SubAnimeNode | SubMangaNode] struct {
	Node               *SubEntityNode `json:"node,omitempty"`
	NumRecommendations *int64         `json:"num_recommendations,omitempty"`
}

type ResponsePaging struct {
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
}

type ResponsePage[Entity Anime | Manga] struct {
	Data *[]struct {
		Node Entity `json:"node,omitempty"`
	} `json:"data,omitempty"`
	Paging *ResponsePaging `json:"paging,omitempty"`
}
