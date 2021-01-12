package store

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"bitbucket.org/icehousecorp/moviedb/core"
)

var (
	ActionNowPlaying = "now_playing"
	ActionPopular    = "popular"
	ActionUpcoming   = "upcoming"
)

type MovieStore struct {
	client  *http.Client
	baseURL string
}

func NewMovieStore(client *http.Client, baseURL string) core.MovieStore {
	return MovieStore{
		client:  client,
		baseURL: baseURL,
	}
}

func (ms MovieStore) Request(ctx context.Context, form *core.Request) (*core.Response, error) {
	movieURL := fmt.Sprintf("%s/movie/%s", ms.baseURL, form.Action)
	query := url.Values{}
	query.Set("api_key", form.ApiKey)
	query.Set("language", form.Language)
	query.Set("page", strconv.Itoa(form.Page))
	query.Set("region", form.Region)

	requestURL, err := url.Parse(movieURL)
	if err != nil {
		return nil, fmt.Errorf("cannot parse URL: %w", err)
	}
	requestURL.RawQuery = query.Encode()

	resp, err := ms.client.Get(requestURL.String())
	if err != nil {
		return nil, fmt.Errorf("err request get: %w", err)
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err read body: %w", err)
	}

	var result core.Response
	if err := json.Unmarshal(respByte, &result); err != nil {
		return nil, fmt.Errorf("err decoding json: %w", err)
	}

	return &result, nil
}
