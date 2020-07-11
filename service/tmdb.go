package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mvindex/model"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

const baseURL = "https://api.themoviedb.org/3"

func GetTopRatedList() {
	var conf model.ApiConfig
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	url := baseURL + "/movie/top_rated"
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("api_key", conf.Key)
	req.URL.RawQuery = q.Encode()
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	t := model.TopRatedBase{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

}
