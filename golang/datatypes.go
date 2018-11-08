package main

type Result struct {
	Result [][]Vote `json:"result"`
}

type Vote struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Voter struct {
	Votes [][]Vote `json:"votes"`
}

type InputList struct {
	Voters []Voter `json:"voters"`
}
