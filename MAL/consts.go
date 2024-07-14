package MyAnimeListSDK

type (
	MalAnimeSeason       = string
	AnimeRankingCategory = string
	MangaRankingCategory = string
)

const (
	MinResultListSize     = 1
	DefaultResultListSize = 10
	MaxResultListSize     = 100
)

const (
	MalAnimeSeasonWinter MalAnimeSeason = "winter"
	MalAnimeSeasonSpring MalAnimeSeason = "spring"
	MalAnimeSeasonSummer MalAnimeSeason = "summer"
	MalAnimeSeasonFall   MalAnimeSeason = "fall"
)

const (
	AnimeRankingAll        AnimeRankingCategory = "all"
	AnimeRankingAiring     AnimeRankingCategory = "airing"
	AnimeRankingUpcoming   AnimeRankingCategory = "upcoming"
	AnimeRankingTv         AnimeRankingCategory = "tv"
	AnimeRankingOva        AnimeRankingCategory = "ova"
	AnimeRankingMovie      AnimeRankingCategory = "movie"
	AnimeRankingSpecial    AnimeRankingCategory = "special"
	AnimeRankingPopularity AnimeRankingCategory = "bypopularity"
	AnimeRankingFavorite   AnimeRankingCategory = "favorite"
)
