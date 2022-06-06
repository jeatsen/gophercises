package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type QuestionRecord struct {
	Question string
	Answer string
}

func createQuestionList(data [][]string) []QuestionRecord {
	var questionList []QuestionRecord
	for i, line := range data {
		if i == 0 {
			continue;
		}
		if i > 0 { // omit header line
			var rec QuestionRecord
			for j, field := range line {
				if j == 0 {
					rec.Question = field
				} else if j == 1 {
					rec.Answer = field
				}
			}
			questionList = append(questionList, rec)
		}
	}
	return questionList
}

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	questionList := createQuestionList(data)

	// print the array
	fmt.Printf("%+v\n", questionList)
}