package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mvindex/model/tmdb"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

const baseURL = "https://api.themoviedb.org/3"
const imageURL = "https://image.tmdb.org/t/p/w500"

func getApiKey() string {
	var conf tmdb.ApiConfig
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conf.Key
}

func GetTopRatedList(page string) []tmdb.BaseContents {
	url := baseURL + "/movie/top_rated"
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("api_key", getApiKey())
	q.Add("page", page)
	req.URL.RawQuery = q.Encode()
	resBody := ExecuteRequest(req)
	t := tmdb.Base{}
	err := json.Unmarshal(resBody, &t)
	if err != nil {
		log.Fatal(err)
	}
	for i, r := range t.Rusults {
		t.Rusults[i].PosterPath = imageURL + r.PosterPath
	}
	result := t.Rusults
	return result
}

func GetDetail(id string) tmdb.MovieDetail {
	url := baseURL + "/movie/" + id
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("api_key", getApiKey())
	req.URL.RawQuery = q.Encode()
	resBody := ExecuteRequest(req)
	detail := tmdb.MovieDetail{}
	err := json.Unmarshal(resBody, &detail)
	if err != nil {
		log.Fatal(err)
	}
	detail.PosterPath = imageURL + detail.PosterPath
	return detail
}

func SearchByKeyword(keyword string) []tmdb.BaseContents {
	url := baseURL + "/search/movie"
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("api_key", getApiKey())
	q.Add("query", keyword)
	req.URL.RawQuery = q.Encode()
	resBody := ExecuteRequest(req)
	contents := tmdb.Base{}
	err := json.Unmarshal(resBody, &contents)
	if err != nil {
		log.Fatal(err)
	}
	for i, c := range contents.Rusults {
		contents.Rusults[i].PosterPath = imageURL + c.PosterPath
	}
	return contents.Rusults
}

func ExecuteRequest(req *http.Request) []byte {
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body
}
