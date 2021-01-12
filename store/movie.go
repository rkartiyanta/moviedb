package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"

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

func (ms MovieStore) Request(form *core.Request) (*core.ResponseAll, error) {
	query := url.Values{}
	query.Set("api_key", form.ApiKey)
	query.Set("language", form.Language)
	query.Set("page", strconv.Itoa(form.Page))
	query.Set("region", form.Region)

	var wg sync.WaitGroup
	requestsChan := make(chan *core.Response, 3)
	wg.Add(1)
	// now playing
	go func() {
		defer wg.Done()
		nowPlayingURL := fmt.Sprintf("%s/movie/now_playing", ms.baseURL)
		requestNowPlayingURL, _ := url.Parse(nowPlayingURL)
		requestNowPlayingURL.RawQuery = query.Encode()

		respNowPlaying, err := ms.nowPlaying(requestNowPlayingURL)
		if err == nil {
			requestsChan <- respNowPlaying
		}
	}()

	wg.Add(1)
	// popular
	go func() {
		defer wg.Done()
		popularURL := fmt.Sprintf("%s/movie/popular", ms.baseURL)
		requestPopularURL, _ := url.Parse(popularURL)
		requestPopularURL.RawQuery = query.Encode()

		respPopular, err := ms.popular(requestPopularURL)
		if err == nil {
			requestsChan <- respPopular
		}
	}()

	wg.Add(1)
	// upcoming
	go func() {
		defer wg.Done()
		upcomingURL := fmt.Sprintf("%s/movie/upcoming", ms.baseURL)
		requestUpcomingURL, _ := url.Parse(upcomingURL)
		requestUpcomingURL.RawQuery = query.Encode()

		respUpcoming, err := ms.upcoming(requestUpcomingURL)
		if err == nil {
			requestsChan <- respUpcoming
		}
	}()

	go func() {
		defer close(requestsChan)
		wg.Wait()
	}()

	result := core.ResponseAll{}
	for value := range requestsChan {
		if value.Action == "now_playing" {
			result.ReponseNowPlaying = *value
		}
		if value.Action == "upcoming" {
			result.ReponseUpcoming = *value
		}
		if value.Action == "popular" {
			result.ReponsePopular = *value
		}
	}

	return &result, nil
}

func (ms MovieStore) nowPlaying(requestURL *url.URL) (*core.Response, error) {
	resp, err := ms.client.Get(requestURL.String())
	if err != nil {
		return nil, fmt.Errorf("err now playing: %w", err)
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err read body: %w", err)
	}

	var result core.Response
	if err := json.Unmarshal(respByte, &result); err != nil {
		return nil, fmt.Errorf("err decoding json: %w", err)
	}

	result.Action = "now_playing"
	return &result, nil
}

func (ms MovieStore) popular(requestURL *url.URL) (*core.Response, error) {
	resp, err := ms.client.Get(requestURL.String())
	if err != nil {
		return nil, fmt.Errorf("err now playing: %w", err)
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err read body: %w", err)
	}

	var result core.Response
	if err := json.Unmarshal(respByte, &result); err != nil {
		return nil, fmt.Errorf("err decoding json: %w", err)
	}

	result.Action = "popular"
	return &result, nil
}

func (ms MovieStore) upcoming(requestURL *url.URL) (*core.Response, error) {
	resp, err := ms.client.Get(requestURL.String())
	if err != nil {
		return nil, fmt.Errorf("err now playing: %w", err)
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err read body: %w", err)
	}

	var result core.Response
	if err := json.Unmarshal(respByte, &result); err != nil {
		return nil, fmt.Errorf("err decoding json: %w", err)
	}

	result.Action = "upcoming"
	return &result, nil
}
