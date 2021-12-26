package params

import (
	"net/http"
	"strconv"

	"log"
)

const (
	defaultItemsPerPage = 25
	defaultPage         = 1
)

type QueryParams struct {
	Search Search
}

type Search struct {
	ItemsPerPage int
	Page         int
}

func New(r *http.Request) QueryParams {
	query := r.URL.Query()

	return QueryParams{
		Search: Search{
			ItemsPerPage: parseInt(query.Get("itemsPerPage"), defaultItemsPerPage),
			Page:         parseInt(query.Get("page"), defaultPage),
		},
	}
}

func parseInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("failed to parse int: %v", err)
		return defaultValue
	}

	return i
}
