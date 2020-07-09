package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://developers.themoviedb.org/3"
const apiKey = "apiKey"

func GetTopRatedList() {
	url := baseURL + "/movie/top_rated?" + apiKey
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	byteArray, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))

}
