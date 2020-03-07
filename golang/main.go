package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

//Argument documentation
//
// -v | --validate		Validates input and output. (What even does this mean?)
// -s | --source		Source file location. Either source file or input is needed for the program to run.
// -i | --input			Input voting data here in a JSON formatted string. Either input or source file is needed for the program to run.
// -d | --destination	Destination file location, if not provided, result will be returned JSON formatted.
// -e | --expected		Expected value. (What even does this mean?)
func main() {
	sourceFlag := flag.String("source", "", "Provide a file path here")
	destinationFlag := flag.String("destination", "", "Provide a file path here")
	// dFlat := flag.String("d", "", "Provide a file path here")	}
	flag.Parse()

	sourceFile := *sourceFlag
	destinationFile := *destinationFlag

	var il NewInputList
	readFile := sourceFile
	fmt.Println("Reading from file " + readFile)
	bytes, e := ioutil.ReadFile(readFile)
	check(e)
	json.Unmarshal(bytes, &il)
	fmt.Println("Finished reading JSON")

	result := calculateNewModel(il, destinationFile)

	resultJSON, e := json.MarshalIndent(result, "", "    ")
	check(e)
	writeFile := destinationFile
	e = ioutil.WriteFile(writeFile, resultJSON, 0644)
	fmt.Println("Writing to file " + writeFile)
	check(e)
}

func calculateNewModel(il NewInputList, destinationFile string) NewResult {
	start := time.Now()

	result, e := calc(il)
	check(e)

	elapsed := time.Now().Sub(start)
	fmt.Println("Operation took " + elapsed.String())
	fmt.Println()
	fmt.Println("===================")
	fmt.Println()

	return result
}

func convertToNew(number int) {
	start := time.Now()
	file := strconv.Itoa(number)
	var il InputList
	readFile := "../../kvalitetsmetoden_testfiles/test" + file + "_in_legacy.json"
	fmt.Println("Reading from file " + readFile)
	bytes, e := ioutil.ReadFile(readFile)
	check(e)
	json.Unmarshal(bytes, &il)
	fmt.Println("Finished reading JSON")

	newList := il.convertToNewJSON()
	resultJSON, e := json.Marshal(newList)
	check(e)
	writeFile := "../../kvalitetsmetoden_testfiles/test" + file + "_in.json"
	e = ioutil.WriteFile(writeFile, resultJSON, 0644)
	fmt.Println("Writing to file " + writeFile)
	check(e)

	elapsed := time.Now().Sub(start)
	fmt.Println("Operation took " + elapsed.String())
	fmt.Println()
	fmt.Println("===================")
	fmt.Println()
}
