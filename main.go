package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)
const 
(
	defaultFilename ="Problems.csv"
	defaultTime =30
)

type question struct{
	question string
    answer string
}

type result struct{
	questionLen int
	answerLen int
	correct int
	incorrect int
}

func main(){	
 

	flagname := flag.String("name", defaultFilename, "File name")
	flagtimer := flag.Int("time", defaultTime, "timer for quiz  default 3s")
	flagShuffle := flag.Bool("s", false, "Shuffle the quiz questions ")

    flag.Parse()
	data := readCSV(*flagname)
	if *flagShuffle{
       shuffle(data)
	}
	quizP2(data, *flagtimer)
} 

//getCSVdata open csv file  by name  and return the data from it
func readCSV(name string)[]question{

	file, err := os.Open(name)
	if err !=nil {
		fmt.Printf("failed to open file: %v\n", err)
		os.Exit(0)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err !=nil {
		fmt.Printf("failed to open file: %v\n", err)
		os.Exit(0)
	}
	q := make([]question, len(records))
	for i := range records{
       q[i]= question{question: records[i][0], answer: records[i][1]}
	}

  return q
}

 
func quizP2(questions []question, timer int){

	ctx, cancle:= context.WithCancel(context.Background())
	var r result
	go  ask(cancle, questions, &r)
	select{
		case <-time.After(time.Second * time.Duration(timer)):
			fmt.Printf("\nTime out \nQuestions: %v \\ %v \nNumber of correct answer : %v \nNumber of incorrect answer: %v\n", r.questionLen, r.answerLen , r.correct, r.incorrect)	
		case <- ctx.Done():
			fmt.Printf("Questions: %v \\ %v \nNumber of correct answer : %v \nNumber of incorrect answer: %v\n", r.questionLen, r.answerLen , r.correct, r.incorrect)	
	}
}

func ask(cancle context.CancelFunc , questions []question, r  *result )  {
	scanner := bufio.NewScanner(os.Stdin)
	var ans string
	r.questionLen = len(questions)
	for i,v:= range questions {

		fmt.Printf("Q.%v %v ",(i+1), v.question)
		scanner.Scan()
	    ans =  strings.TrimSpace(scanner.Text())
		r.answerLen ++
		if ans == strings.TrimSpace(v.answer){
			r.correct ++
			continue
		}
 
		r.incorrect ++
		
	
	}
	cancle()
}

func shuffle(s []question){
	rand.Seed(time.Now().UnixNano())
   rand.Shuffle(len(s), func(i, j int ){ s[i], s[j] = s[j], s[i]})
}