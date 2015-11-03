package tmdb

import (
	"fmt"
	"github.com/jmcvetta/napping"
)

type Tmdb struct {
	ApiKey        string
	api           napping.Session
	BaseUrl       string
	SecureBaseUrl string
}

const (
	configuration = "https://api.themoviedb.org/3/configuration"
	searchmovie   = "https://api.themoviedb.org/3/search/movie"
	movie         = "https://api.themoviedb.org/3/movie/%d"
)

type ConfigurationResult struct {
	Images struct {
		Base_Url        string
		Secure_Base_Url string
		Poster_Sizes    []string
		Backdrop_Sizes  []string
		Profile_Sizes   []string
		Logo_Sizes      []string
	}
	ChangeKeys []string
}

type SearchMovieResult struct {
	Adult          bool    `json: adult`
	Backdrop_Path  string  `json: backdrop_path`
	Id             uint64  `json: id`
	Original_Title string  `json: original_title`
	Release_Date   string  `json: release_date`
	Poster_Path    string  `json: poster_path`
	Popularity     float64 `json: popularity`
	Title          string  `json: title`
	Vote_Average   float64 `json: vote_average`
	Vote_Count     uint64  `json: vote_count`
}

type SearchMovieResponse struct {
	Page          uint64              `json: page`
	Results       []SearchMovieResult `json: results`
	Total_Pages   uint64              `json: total_pages`
	Total_Results uint64              `json: total_results`
}

type Collection struct {
	Id            uint64 `json: id`
	Name          string `json: name`
	Poster_Path   string `json: poster_path`
	Backdrop_Path string `json: backdrop_path`
}

type Genre struct {
	Id   uint64 `json: id`
	Name string `json: name`
}

type Company struct {
	Name string `json: name`
	Id   uint64 `json: id`
}

type Country struct {
	ISO_3166_1 string `json: iso_3166_1`
	Name       string `json: name`
}

type Language struct {
	ISO_639_1 string `json: iso_639_1`
	Name      string `json: name`
}

type GetMovieResponse struct {
	Adult                 bool       `json: adult`
	Backdrop_Path         string     `json: backdrop_path`
	Belongs_To_Collection Collection `json: belongs_to_collection`
	Budget                uint64     `json: budget`
	Genres                []Genre    `json: genres`
	Homepage              string     `json: homepage`
	Id                    uint64     `json: id`
	Imdb_Id               string     `json: imdb_id`
	Original_Title        string     `json: original_title`
	Overview              string     `json: overview`
	Popularity            float64    `json: popularity`
	Poster_Path           string     `json: poster_path`
	Production_Companies  []Company  `json: production_companies`
	Production_Countries  []Country  `json: production_countries`
	Release_Date          string     `json: release_date`
	Revenue               uint64     `json: revenue`
	Runtime               uint64     `json: runtime`
	Spoken_Languages      []Language `json: spoken_languages`
	Status                string     `json: status`
	Tagline               string     `json: tagline`
	Title                 string     `json: title`
	Vote_Average          float64    `json: vote_average`
	Vote_Count            uint64     `json: vote_count`
}

func NewClient(apiKey string, log bool) (tmdb *Tmdb, err error) {
	session := napping.Session{Log: log}

	res := ConfigurationResult{}
	params := napping.Params{
		"api_key": apiKey,
	}.AsUrlValues()

	resp, err := session.Get(configuration, &params, &res, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to get TMDB Configuration: %s", err)
	}

	if resp.Status() != 200 {
		return nil, fmt.Errorf("Error getting TMDB configuration: [status %v] %v", resp.Status(), resp.RawText())
	}

	return &Tmdb{apiKey, session, res.Images.Base_Url, res.Images.Secure_Base_Url}, nil
}

func (self *Tmdb) SearchMovie(title string) (res *SearchMovieResponse, err error) {
	params := napping.Params{
		"api_key": self.ApiKey,
		"query":   title,
	}.AsUrlValues()

	resp, err := self.api.Get(searchmovie, &params, &res, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to searchMovie %s", err)
	}

	if resp.Status() != 200 {
		return nil, fmt.Errorf("Bad response from SearchMovie: [status %v] %v", resp.Status(), resp.RawText())
	}

	return
}

func (self *Tmdb) GetMovie(id uint64) (res *GetMovieResponse, err error) {
	params := napping.Params{
		"api_key": self.ApiKey,
	}.AsUrlValues()

	resp, err := self.api.Get(fmt.Sprintf(movie, id), &params, &res, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to getMovie %s", err)
	}

	if resp.Status() != 200 {
		return nil, fmt.Errorf("Bad response from GetMovie: [status %v] %v", resp.Status(), resp.RawText())
	}

	return
}
