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
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	t := tmdb.Base{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Fatal(err)
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
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	detail := tmdb.MovieDetail{}
	err = json.Unmarshal(body, &detail)
	if err != nil {
		log.Fatal(err)
	}
	return detail
}

func SearchByKeyword(keyword string) []tmdb.BaseContents {
	url := baseURL + "/search/movie"
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("api_key", getApiKey())
	q.Add("query", keyword)
	req.URL.RawQuery = q.Encode()
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	contents := tmdb.Base{}
	err = json.Unmarshal(body, &contents)
	return contents.Rusults
}
