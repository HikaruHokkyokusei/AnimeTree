package MALModels

import (
	"reflect"
	"sort"
	"strings"
)

var (
	allAnimeFields string
)

func fieldsToDelimitedJsonListStr(structure any, delimiter string, patcher func(string) string) string {
	if delimiter == "" {
		delimiter = ","
	}
	var res []string
	val := reflect.ValueOf(structure)
	for i := 0; i < val.Type().NumField(); i++ {
		jsonName := strings.ReplaceAll(val.Type().Field(i).Tag.Get("json"), ",omitempty", "")
		if patcher != nil {
			jsonName = patcher(jsonName)
		}
		res = append(res, jsonName)
	}
	sort.Strings(res)
	return strings.Join(res, delimiter)
}

func getSubAnimeFields() string {
	return fieldsToDelimitedJsonListStr(SubAnimeNode{}, ",", func(jsonName string) string {
		if jsonName == "my_list_status" {
			return jsonName + "{finish_date,is_rewatching,num_episodes_watched,score,start_date,status,updated_at}"
		}
		return jsonName
	})
}

func getAllAnimeFields() string {
	tier1AnimeFields := getSubAnimeFields()
	tier2AnimeFields := fieldsToDelimitedJsonListStr(Anime{}, ",", func(jsonName string) string {
		if jsonName == "recommendations" || jsonName == "related_anime" {
			return jsonName + "{" + tier1AnimeFields + "}"
		}
		return jsonName
	})
	return tier2AnimeFields
}

func init() {
	allAnimeFields = getAllAnimeFields()
}

func AllAnimeFields() string {
	return allAnimeFields
}
