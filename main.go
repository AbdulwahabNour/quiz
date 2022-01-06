package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

//add a timer. The  time limit should be 30 seconds,
//customiz timer by flag
//Quiz should stop as soon as the time limit has execeeded
//Users should be asked to press enter (or some other key)
 
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
	correct int
	incorrect int
}

func main(){	
 

	name := flag.String("name", defaultFilename, "File name")
	timer := flag.Int("time", defaultTime, "timer for quiz  default 3s")
    flag.Parse()
	data := readCSV(*name)
	quizP2(data, *timer)
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
			fmt.Printf("\nTime out \nNumber of questions: %v \nNumber of correct answer : %v \nNumber of incorrect answer: %v\n", r.questionLen , r.correct, r.incorrect)	
		case <- ctx.Done():
			fmt.Printf("\nNumber of questions: %v \nNumber of correct answer : %v \nNumber of incorrect answer: %v\n", r.questionLen, r.correct, r.incorrect)	
	}
}

func ask(cancle context.CancelFunc , questions []question, r  *result )  {
	scanner := bufio.NewScanner(os.Stdin)
	var ans string
	r.questionLen = len(questions)
	for i,v:= range questions {

		fmt.Printf("Q.%v %v ",(i+1), v.question)
		scanner.Scan()
	    ans = scanner.Text()
		if ans == v.answer{
			r.correct ++
			continue
		}
 
		r.incorrect ++
	
	}
	cancle()
}