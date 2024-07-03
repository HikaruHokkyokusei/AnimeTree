package MALModels

import (
	"reflect"
	"sort"
	"strings"
)

var (
	allAnimeFields string
	allMangaFields string
)

func fieldsToDelimitedJsonListStr(structure any, delimiter string) string {
	if delimiter == "" {
		delimiter = ","
	}
	var res []string
	val := reflect.ValueOf(structure)
	for i := 0; i < val.Type().NumField(); i++ {
		jsonName := strings.ReplaceAll(val.Type().Field(i).Tag.Get("json"), ",omitempty", "")
		res = append(res, jsonName)
	}
	sort.Strings(res)
	return strings.Join(res, delimiter)
}

func getAllEntityFields[
	Entity Anime | Manga,
	EntityMyListStatus AnimeMyListStatus | MangaMyListStatus,
	SubEntityNode SubAnimeNode | SubMangaNode,
](entity Entity, entityMyListStatus EntityMyListStatus, subEntityNode SubEntityNode) string {
	myListStatusFields := fieldsToDelimitedJsonListStr(entityMyListStatus, ",")
	recommendationsFields := fieldsToDelimitedJsonListStr(subEntityNode, ",")

	subAnimeFields := fieldsToDelimitedJsonListStr(SubAnimeNode{}, ",")
	subMangaFields := fieldsToDelimitedJsonListStr(SubMangaNode{}, ",")

	entityBaseFields := fieldsToDelimitedJsonListStr(entity, ",")

	entityBaseFields = strings.ReplaceAll(
		entityBaseFields, "my_list_status",
		"my_list_status{"+myListStatusFields+"}",
	)
	entityBaseFields = strings.ReplaceAll(
		entityBaseFields, "recommendations",
		"recommendations{"+recommendationsFields+"}",
	)
	entityBaseFields = strings.ReplaceAll(
		entityBaseFields, "related_anime",
		"related_anime{"+subAnimeFields+"}",
	)
	entityBaseFields = strings.ReplaceAll(
		entityBaseFields, "related_manga",
		"related_manga{"+subMangaFields+"}",
	)

	return entityBaseFields
}

func init() {
	allAnimeFields = getAllEntityFields(Anime{}, AnimeMyListStatus{}, SubAnimeNode{})
	allMangaFields = getAllEntityFields(Manga{}, MangaMyListStatus{}, SubMangaNode{})
}

func AllAnimeFields() string {
	return allAnimeFields
}

func AllMangaFields() string {
	return allMangaFields
}
