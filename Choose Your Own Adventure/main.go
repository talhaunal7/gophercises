package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Game map[string]Story

func main() {
	jsonFile, err := ioutil.ReadFile("/Users/talhaunal/Programming/go projects/gophercises/Choose Your Own Adventure/gopher.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data Game
	json.Unmarshal(jsonFile, &data)
	fmt.Println(data["intro"].Title)
	for _, v := range data {
		fmt.Println(v.Title)
	}

}
