package pagination

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ParseGetListPageToken(pageToken string) (lastValue uint64) {
	pToken := strings.Split(pageToken, "_")

	if len(pToken) != 2 {
		return
	}

	sortValue := pToken[1]
	lastValue, _ = strconv.ParseUint(sortValue, 10, 64)

	return
}

func CreateGetListPageToken(data interface{}, limit int) (nextToken string) {
	voData := reflect.ValueOf(data)
	if voData.Kind() != reflect.Slice {
		return
	}

	if !(voData.IsValid()) {
		return
	}

	if voData.Len() < limit {
		return
	}

	lastData := voData.Index(voData.Len() - 1)
	nextTimestamp := time.Now().Unix()

	nextID := lastData.FieldByName("UpdatedAt")
	nextToken = fmt.Sprintf("%d_%d", nextTimestamp, nextID)

	return
}
