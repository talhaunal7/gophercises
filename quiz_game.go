package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var correctCount int

func getTotalQuestionCount() (int, error) {
	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1, errors.New("error opening the file")
	}

	defer file.Close()
	reader := csv.NewReader(file)
	totalQuestions := 0

	for {
		_, err := reader.Read()

		if err != io.EOF {
			totalQuestions++
		} else {
			break
		}
	}

	return totalQuestions, nil
}

func getInput(done chan bool) {
	problemNumber := 1
	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()
	reader := csv.NewReader(file)

	for {
		var input, data int
		record, err := reader.Read()
		if err == io.EOF {
			done <- true
			break
		}
		if err != nil {
			fmt.Println("Error reading record:", err)
			return
		}
		fmt.Println("Problem#" + fmt.Sprint(problemNumber) + " " + record[0] + " = ?")
		problemNumber++
		_, er := fmt.Scanf("%d", &input)
		if er != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		data, _ = strconv.Atoi(record[1])
		if data == input {
			correctCount++
		}
	}

}

func main() {
	done := make(chan bool)
	totalQ, _ := getTotalQuestionCount()
	numbPtr := flag.Int("limit", 10, "an int")
	flag.Parse()
	fmt.Println("Press Enter to Start the Game")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	go getInput(done)

	select {
	case <-done:
		fmt.Println("end")
	case <-time.After(time.Second * time.Duration(*numbPtr)):
		fmt.Println("timed out")
	}

	fmt.Println("You scored " + fmt.Sprint(correctCount) + " out of " + fmt.Sprint(totalQ))
}
