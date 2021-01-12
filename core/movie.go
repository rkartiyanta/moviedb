package core

import (
	"context"
)

type (
	Request struct {
		Action   string
		ApiKey   string
		Language string
		Page     int
		Region   string
	}

	Response struct {
		Page         int              `json:"page"`
		Results      []ResponseResult `json:"results"`
		Dates        ResponseDates    `json:"dates"`
		TotalPages   int              `json:"total_pages"`
		TotalResults int              `json:"total_results"`
	}

	ResponseResult struct {
		PosterPath       *string `json:"poster_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLangauge string  `json:"original_Language"`
		Title            string  `json:"title"`
		BackdropPath     *string `json:"backdrop_path"`
		Popularity       float32 `json:"popularity"`
		Video            bool    `json:"video"`
		VoteCount        int     `json:"vote_count"`
		VoteAverage      float32 `json:"vote_average"`
	}

	ResponseDates struct {
		Max string `json:"maximum"`
		Min string `json:"minimum"`
	}

	MovieStore interface {
		Request(ctx context.Context, form *Request) (*Response, error)
	}
)
