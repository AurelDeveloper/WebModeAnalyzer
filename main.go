package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type UrlList struct {
	Urls []string `json:"urls"`
}

type PageMode struct {
	Url  string `json:"url"`
	Mode string `json:"mode"`
}

func checkMode(content string) string {
	if strings.Contains(content, "#000") || strings.Contains(content, "rgb(0, 0, 0)") {
		return "Dark Mode"
	}
	return "Light Mode"
}

func main() {
	data, err := ioutil.ReadFile("urls.json")
	if err != nil {
		log.Fatal(err)
	}

	var urlList UrlList
	err = json.Unmarshal(data, &urlList)
	if err != nil {
		log.Fatal(err)
	}

	var pageModes []PageMode

	for _, url := range urlList.Urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		mode := checkMode(string(body))
		pageModes = append(pageModes, PageMode{Url: url, Mode: mode})
	}

	output, err := json.Marshal(pageModes)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("output.json", output, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
