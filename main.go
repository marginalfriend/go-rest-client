package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Album struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/albums"
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	bodyString := string(bodyBytes)

	var albums []Album

	err = json.Unmarshal([]byte(bodyString), &albums)

	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		swapped := false

		for i := 1; i < len(albums); i++ {
			if albums[i].Id > albums[i-1].Id {
				bucket := albums[i].Id
				albums[i].Id = albums[i-1].Id
				albums[i-1].Id = bucket
				swapped = true
			}

		}

		if !swapped {
			break
		}
	}

	for _, album := range albums {
		fmt.Println(album.Id)
	}
}
