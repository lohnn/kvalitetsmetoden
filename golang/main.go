package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var il InputList
	bytes, e := ioutil.ReadFile("test_files/test1_in.json")
	check(e)
	json.Unmarshal(bytes, &il)
	fmt.Println(il)

	result, e := calc(il)
	check(e)
	resultJson, e := json.Marshal(result)
	check(e)
	fmt.Println(string(resultJson))
}