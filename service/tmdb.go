package service

import (
	"fmt"
	"log"
	"mvindex/model"

	"github.com/kelseyhightower/envconfig"
)

const baseURL = "https://developers.themoviedb.org/3"

func GetTopRatedList() {
	var conf model.ApiConfig
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(conf.Key)
	// url := baseURL + "/movie/top_rated?" + apiKey
	// req, _ := http.NewRequest("GET", url, nil)
	// client := new(http.Client)
	// res, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// byteArray, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(byteArray))

}
