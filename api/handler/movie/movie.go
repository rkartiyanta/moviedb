package movie

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"bitbucket.org/icehousecorp/moviedb/api/handler"
	"bitbucket.org/icehousecorp/moviedb/core"
	pkgerr "bitbucket.org/icehousecorp/moviedb/pkg/error"
)

func Request(movieStore core.MovieStore) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		requestParam, err := populateRequest(r.URL.Query())
		if err != nil {
			handler.Write(w, err)
			return
		}

		result, err := movieStore.Request(requestParam)
		if err != nil {
			handler.Write(w, err)
			return
		}

		handler.Write(w, result)
	}
	return handler
}

func populateRequest(query url.Values) (*core.Request, error) {
	page, err := strconv.Atoi(query.Get("page"))
	if page < 1 || err != nil {
		page = 1
	}

	apiKey := strings.TrimSpace(query.Get("api_key"))
	if apiKey == "" {
		return nil, pkgerr.UnauthorizeError{
			StatusMessage: "Invalid API key: You must be granted a valid key.",
			Success:       false,
			StatusCode:    7,
		}
	}

	return &core.Request{
		ApiKey:   apiKey,
		Language: strings.TrimSpace(query.Get("language")),
		Page:     page,
		Region:   strings.TrimSpace(query.Get("region")),
	}, nil
}
