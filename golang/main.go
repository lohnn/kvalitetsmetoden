package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	convertToNew(1)
}

func calculateNewModel(number int) {
	var start = time.Now()
	var file = strconv.Itoa(number)
	var il NewInputList
	var readFile = "test_files/test" + file + "_in.json"
	fmt.Println("Reading from file " + readFile)
	bytes, e := ioutil.ReadFile(readFile)
	check(e)
	json.Unmarshal(bytes, &il)
	fmt.Println("Finished reading JSON")

	result, e := calc(il)
	check(e)
	resultJson, e := json.Marshal(result)
	check(e)
	var writeFile = "test_files/test" + file + "_out.json"
	e = ioutil.WriteFile(writeFile, resultJson, 0644)
	fmt.Println("Writing to file " + writeFile)
	check(e)

	var elapsed = time.Now().Sub(start)
	fmt.Println("Operation took " + elapsed.String())
	fmt.Println()
	fmt.Println("===================")
	fmt.Println()
}

func convertToNew(number int) {
	var start = time.Now()
	var file = strconv.Itoa(number)
	var il InputList
	var readFile = "test_files/test" + file + "_in_legacy.json"
	fmt.Println("Reading from file " + readFile)
	bytes, e := ioutil.ReadFile(readFile)
	check(e)
	json.Unmarshal(bytes, &il)
	fmt.Println("Finished reading JSON")

	newList := il.convertToNewJson()
	resultJson, e := json.Marshal(newList)
	check(e)
	var writeFile = "test_files/test" + file + "_in.json"
	e = ioutil.WriteFile(writeFile, resultJson, 0644)
	fmt.Println("Writing to file " + writeFile)
	check(e)

	var elapsed = time.Now().Sub(start)
	fmt.Println("Operation took " + elapsed.String())
	fmt.Println()
	fmt.Println("===================")
	fmt.Println()
}
