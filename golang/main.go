package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	calculateNewModel(1)
}

func calculateNewModel(number int) {
	start := time.Now()
	file := strconv.Itoa(number)
	var il NewInputList
	readFile := "../../kvalitetsmetoden_testfiles/test" + file + "_in.json"
	fmt.Println("Reading from file " + readFile)
	bytes, e := ioutil.ReadFile(readFile)
	check(e)
	json.Unmarshal(bytes, &il)
	fmt.Println("Finished reading JSON")

	result, e := calc(il)
	check(e)
	resultJSON, e := json.Marshal(result)
	check(e)
	writeFile := "../../kvalitetsmetoden_testfiles/test" + file + "_out.json"
	e = ioutil.WriteFile(writeFile, resultJSON, 0644)
	fmt.Println("Writing to file " + writeFile)
	check(e)

	elapsed := time.Now().Sub(start)
	fmt.Println("Operation took " + elapsed.String())
	fmt.Println()
	fmt.Println("===================")
	fmt.Println()
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

	newList := il.convertToNewJson()
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
