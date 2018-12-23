package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)

	}

	fmt.Printf("%s\n", data)

	data1, _ := json.MarshalIndent(movies, "*", "----")
	fmt.Printf("%s\n", data1)

	var titles []struct{Title string
						Year int `json:"released"`}

	if err := json.Unmarshal(data, &titles); err !=nil {
		log.Fatalf("JSON unmarshaling failed: %s\n", err)
	}

	fmt.Printf("%v\n",titles)
	
	//var ii int = 16
	//fmt.Println(json.Marshal(ii))
}