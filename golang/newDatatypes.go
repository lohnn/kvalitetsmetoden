package main

type newVote [][]int

func (votesList newVote) flatten() []int {
	var votes, voteArray []int
	for _, voteArray = range votesList {
		votes = append(votes, voteArray...)
	}
	return votes
}

func flatten(votesList newVote) []int {
	var votes, voteArray []int
	for _, voteArray = range votesList {
		votes = append(votes, voteArray...)
	}
	return votes
}

func (votesList newVote) mapToVotes(voteList VoteList) [][]Vote {
	var returnList [][]Vote

	for _, outer := range votesList {
		returnList = append(returnList, mapToVotes(outer, voteList))
	}

	return returnList
}

func mapToVotes(outer []int, voteList VoteList) []Vote {
	var returnList []Vote

	for _, voteIndex := range outer {
		returnList = append(returnList, voteList[voteIndex])
	}

	return returnList
}

// NewVoter is the new format for a voter
type NewVoter struct {
	Votes newVote `json:"vote"`
}

// NewInputList is the new format of the input list
type NewInputList struct {
	Candidates VoteList  `json:"candidates"`
	Votes      [][][]int `json:"votes"`
}

// NewInputListJSON is the new format of the input list
type NewInputListJSON struct {
	Candidates VoteList  `json:"candidates"`
	Votes      [][][]int `json:"votes"`
}

// NewResult is the new format of the result
type NewResult struct {
	Candidates VoteList `json:"candidates"`
	Result     [][]int  `json:"result"`
}

func (list InputList) convertToNewJSON() NewInputListJSON {
	temp := list.Voters[0].flattenVotes()
	voteMap := temp.mapVotes()

	newVoters := make([][][]int, len(list.Voters))
	for voterIndex, voter := range list.Voters {
		newVotes := make([][]int, len(voter.Votes))
		for voteIndex, vote := range voter.Votes {
			newInnerVotes := make([]int, len(vote))
			for candidateIndex, candidate := range vote {
				newInnerVotes[candidateIndex] = voteMap[candidate]
			}
			newVotes[voteIndex] = newInnerVotes
		}
		newVoters[voterIndex] = newVotes
	}

	inverseMap := make(map[int]Vote)
	for k, v := range voteMap {
		inverseMap[v] = k
	}

	return NewInputListJSON{
		Candidates: temp,
		Votes:      newVoters,
	}
}

func (list InputList) convertToNew() NewInputList {
	temp := list.Voters[0].flattenVotes()
	voteMap := temp.mapVotes()

	newVoters := make([][][]int, len(list.Voters))
	for voterIndex, voter := range list.Voters {
		newVotes := make([][]int, len(voter.Votes))
		for voteIndex, vote := range voter.Votes {
			newInnerVotes := make([]int, len(vote))
			for candidateIndex, candidate := range vote {
				newInnerVotes[candidateIndex] = voteMap[candidate]
			}
			newVotes[voteIndex] = newInnerVotes
		}
		newVoters[voterIndex] = newVotes
	}

	inverseMap := make(map[int]Vote)
	for k, v := range voteMap {
		inverseMap[v] = k
	}

	return NewInputList{
		Candidates: temp,
		Votes:      newVoters,
		//InverseMap: inverseMap,
	}
}

func (voter VoteList) mapVotes() map[Vote]int {
	myMap := make(map[Vote]int)
	for i, vote := range voter {
		myMap[vote] = i
	}
	return myMap
}
