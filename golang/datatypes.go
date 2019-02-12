package main

type Result struct {
	Result Votes `json:"result"`
}

type Vote struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Voter struct {
	Votes Votes `json:"votes"`
}

type Votes [][]Vote

type InputList struct {
	Voters []Voter `json:"voters"`
}

func (votesList Votes) flatten() VoteList {
	var votes, voteArray []Vote
	for _, voteArray = range votesList {
		votes = append(votes, voteArray...)
	}
	return votes
}

func (voter Voter) flattenVotes() VoteList {
	return voter.Votes.flatten()
}

type VoteList []Vote
