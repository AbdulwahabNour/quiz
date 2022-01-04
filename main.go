package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

 
const defaultFilename ="Problems.csv"

func main(){
    var name = flag.String("name", defaultFilename, "File name")
    flag.Parse()
	data := getCSVdata(*name)
	quiz(data)


}

//getCSVdata open csv file  by name  and return the data from it
func getCSVdata(name string)[][]string{
 
   file, err := os.Open(name)
   if err !=nil {
	  fmt.Printf("failed to open file: %v\n", err)
	  os.Exit(0)
   }
	defer file.Close()
	recorsd, err := csv.NewReader(file).ReadAll()
	if err !=nil {
		fmt.Printf("failed to open file: %v\n", err)
		os.Exit(0)
	}
  return recorsd
}

func quiz(question [][]string){
    questionsLen := len(question)
	var right, incorrect uint32
	var ans string

	for i := 0; i < questionsLen; i++ {
		fmt.Printf("Q.%v %v ",(i+1), question[i][0])
		fmt.Scan(&ans)
		if ans == question[i][1]{
			right ++
			continue
		}
 
		incorrect ++
	}

	fmt.Printf("Number of questions: %v \nNumber of right answer : %v \nNumber of incorrect answer: %v\n", questionsLen, right, incorrect)	
}