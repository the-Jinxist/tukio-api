package pkg

import (
	"net/url"
	"strconv"
)

func GetQueryParams(q url.Values) QueryParams {
	limitStr, cursor := q.Get("limit"), q.Get("cursor")

	limit, _ := strconv.Atoi(limitStr)
	if limit == 0 {
		limit = 20
	}

	params := QueryParams{
		Limit:  limit,
		Cursor: cursor,
	}
	return params
}
