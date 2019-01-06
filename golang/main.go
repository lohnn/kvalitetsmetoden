package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	var start = time.Now()
	var file = strconv.Itoa(1)
	var il InputList
	bytes, e := ioutil.ReadFile("test_files/test" + file + "_in.json")
	check(e)
	json.Unmarshal(bytes, &il)
	fmt.Println("Finished reading JSON")

	result, e := calc(il)
	check(e)
	resultJson, e := json.Marshal(result)
	check(e)
	e = ioutil.WriteFile("test_files/test"+file+"_out.json", resultJson, 0644)
	check(e)
	var elapsed = time.Now().Sub(start)
	fmt.Println("Operation took " + elapsed.String())
	fmt.Println()
	fmt.Println("===================")
	fmt.Println()
}