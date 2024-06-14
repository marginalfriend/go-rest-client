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
	// GET request with http
	res, err := http.Get("https://jsonplaceholder.typicode.com/albums")

	// Error handling
	if err != nil {
		log.Fatal(err)
	}

	// This function will be called once the func main() finished its tasks
	// Consider it as an act of cleaning
	defer res.Body.Close()

	// Saving the response to memory (RAM) as bytes
	bodyBytes, err := io.ReadAll(res.Body)

	// Converting bytes to string
	bodyString := string(bodyBytes)

	// Container which will hold unmarshalled data
	// Unmarshalling is converting JSON object to Go object (struct)
	var albums []Album

	// Unmarshalling method
	err = json.Unmarshal([]byte(bodyString), &albums)

	// Error handler
	if err != nil {
		log.Fatal(err)
		return
	}

	// Infinite loop
	for {
		// If swapped, we continue the infinite loop, if not, we break the loop
		swapped := false

		// Iterating through the array of album
		for i := 1; i < len(albums); i++ {

			// This kinda swap (index > index-1) will sort from biggest to smallest
			// Change it to < if you want to sort from smallest to biggest
			if albums[i].Id > albums[i-1].Id {
				// Swapping
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

	// Printing
	for _, album := range albums {
		fmt.Println(album.Id)
	}
}
