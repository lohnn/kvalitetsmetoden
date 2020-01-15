package main

// Result is the data structure that is returned when the calculator is done
type Result struct {
	Result Votes `json:"result"`
}

// Vote is a single vote data structure
type Vote struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// Voter is the data structure for a single voter
type Voter struct {
	Votes Votes `json:"votes"`
}

// Votes is a list of votes
type Votes [][]Vote

// InputList is the format that is inputted to this program
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

// VoteList is a list of votes
type VoteList []Vote
