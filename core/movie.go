package core

type (
	Request struct {
		ApiKey   string
		Language string
		Page     int
		Region   string
	}

	Response struct {
		Action       string           `json:"-"`
		Page         int              `json:"page"`
		Results      []ResponseResult `json:"results"`
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

	ResponseAll struct {
		ReponseNowPlaying Response `json:"now_playing"`
		ReponsePopular    Response `json:"popular"`
		ReponseUpcoming   Response `json:"upcoming"`
	}

	MovieStore interface {
		Request(form *Request) (*ResponseAll, error)
	}
)
