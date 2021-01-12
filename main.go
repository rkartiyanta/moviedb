package main

import (
	"bitbucket.org/icehousecorp/moviedb/api"
	"bitbucket.org/icehousecorp/moviedb/store"
)

func main() {
	movieStore := store.NewMovieStore(store.NewClient(), "https://api.themoviedb.org/3")
	apiServer := api.New(movieStore)
	apiServer.Run()
}
