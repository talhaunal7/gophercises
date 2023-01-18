package main

import (
	"encoding/json"
	"errors"
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

func displayStory(story string, data Game) {
	fmt.Println("---------------------------------")
	fmt.Println("TITLE : ", data[story].Title)
	fmt.Println("STORY : ", data[story].Story)
	fmt.Println("OPTIONS :")
	for i, v := range data[story].Options {
		fmt.Println(i, "-->", v.Text)
	}
	fmt.Println("5 --> Exit Game")
	fmt.Println("---------------------------------")
}

func evaluateInput(input int, optionsCount int) error {
	if input == 5 {
		return errors.New("thanks for playing")
	}

	if input > optionsCount-1 || input < 0 {
		return errors.New("incorrect option choose")
	}
	return nil
}

func main() {
	jsonFile, err := ioutil.ReadFile("/Users/talhaunal/Programming/go projects/go-playground/cyoa/gopher.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var data Game
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		return
	}

	displayStory("intro", data)
	var input int
	optionsCount := len(data["intro"].Options)
	_, er := fmt.Scanf("%d", &input)
	if er != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	er = evaluateInput(input, optionsCount)
	if er != nil {
		fmt.Println(er)
		return
	}

	chosenArc := data["intro"].Options[input].Arc
	optionsCount = len(data[chosenArc].Options)
	for {
		displayStory(chosenArc, data)
		if optionsCount == 0 {
			fmt.Println("THE END")
			break
		}
		_, er = fmt.Scanf("%d", &input)
		if er != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		er = evaluateInput(input, optionsCount)
		if er != nil {
			fmt.Println(er)
			break
		}
		chosenArc = data[chosenArc].Options[input].Arc
		optionsCount = len(data[chosenArc].Options)

	}

}
