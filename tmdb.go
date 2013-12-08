package tmdb

import (
	"github.com/jmcvetta/napping"
	"log"
)

type Tmdb struct {
	ApiKey string
	Api    napping.Session
}

const (
	searchmovie = "https://api.themoviedb.org/3/search/movie"
)

type SearchMovieResult struct {
	Adult bool `json: adult`,
	Backdrop_Path string `json: backdrop_path`,
	Id uint64 `json: id`,
	Original_Title string `json: original_title`,
	Release_Date string `json: release_date`,
	Poster_Path string `json: poster_path`,
	Popularity float64 `json: popularity`,
	Title string `json: title`,
	Vote_Average float64 `json: vote_average`,
	Vote_Count uint64 `json: vote_count`,
}

type SearchMovieResponse struct {
	Page uint64 `json: page`,
	Results []SearchMovieResult `json: results`,
	Total_Pages uint64 `json: total_pages`,
	Total_Results uint64 `json: total_results`,
}

func NewClient(apiKey string) Tmdb {
	return Tmdb{apiKey, napping.Session{}}
}

func (self *Tmdb) SearchMovie(title string) (res SearchMovieResult) {
	params := struct {
		Api_Key string `json: api_key`
	}{
		Api_Key: self.ApiKey
	}

	resp, err := self.Api.Get(seachmovie, &params, &res, nil)
	if err != nil {
		log.Printf("Unable to SearchMovie: %s", err)
		return
	}


}
