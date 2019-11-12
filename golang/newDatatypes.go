package main

type NewVote [][]int

func (votesList NewVote) flatten() []int {
	var votes, voteArray []int
	for _, voteArray = range votesList {
		votes = append(votes, voteArray...)
	}
	return votes
}

func (votesList NewVote) mapToVotes(voteList VoteList) [][]Vote {
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

type NewVoter struct {
	Votes NewVote `json:"vote"`
}

type NewInputList struct {
	//Candidates in the voting, zero indexed, referenced as indexes in Votes
	Candidates VoteList `json:"candidates"`
	//Voting, list of list, uses indexes for referencing candidates
	Votes      []NewVoter   `json:"votes"`
	InverseMap map[int]Vote `json:"-"`
}

type NewInputListJson struct {
	Candidates VoteList		`json:"candidates"`
	Votes      [][][]int	`json:"votes"`
}

func (list InputList) convertToNewJson() NewInputListJson {
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

	return NewInputListJson{
		Candidates: temp,
		Votes:      newVoters,
	}
}


func (list InputList) convertToNew() NewInputList {
	temp := list.Voters[0].flattenVotes()
	voteMap := temp.mapVotes()

	newVoters := make([]NewVoter, len(list.Voters))
	for voterIndex, voter := range list.Voters {
		newVotes := make([][]int, len(voter.Votes))
		for voteIndex, vote := range voter.Votes {
			newInnerVotes := make([]int, len(vote))
			for candidateIndex, candidate := range vote {
				newInnerVotes[candidateIndex] = voteMap[candidate]
			}
			newVotes[voteIndex] = newInnerVotes
		}
		newVoters[voterIndex].Votes = newVotes
	}

	inverseMap := make(map[int]Vote)
	for k, v := range voteMap {
		inverseMap[v] = k
	}

	return NewInputList{
		Candidates: temp,
		Votes:		newVoters,
		InverseMap: inverseMap,
	}
}

func (voter VoteList) mapVotes() map[Vote]int {
	myMap := make(map[Vote]int)
	for i, vote := range voter {
		myMap[vote] = i
	}
	return myMap
}
