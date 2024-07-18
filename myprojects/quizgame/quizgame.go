package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	Question string
	Answer string
}

var problems []Problem
var total int
var numberQuizzes int

func main() {
	problems := readFile()
	var userAnswer string
	
	// Set a timer
	timer := time.AfterFunc(time.Second * 30, func(){
		fmt.Println("You have exhausted the time")
		fmt.Printf("You final score is: %v/%v\n", total, numberQuizzes)	
		os.Exit(0)
	})
	defer timer.Stop()

	for _, problem := range(problems){
		
		fmt.Printf("Question: %v = ", problem.Question)
		fmt.Scan(&userAnswer)
	
		fmt.Println()
		if userAnswer == problem.Answer{
			total += 1
		}
	}
	fmt.Printf("Your score %v/%v", total, numberQuizzes)
}

func readFile() []Problem {
	file, err := os.Open("./problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range(data) {
		problem := Problem {
			Question: row[0],
			Answer: row[1],
		}
		problems = append(problems, problem)		
	}
	numberQuizzes = len(problems)
	return problems
}

