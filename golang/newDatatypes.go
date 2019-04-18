package main

type NewVotes [][]int

func (votesList NewVotes) flatten() []int {
	var votes, voteArray []int
	for _, voteArray = range votesList {
		votes = append(votes, voteArray...)
	}
	return votes
}

func (votesList NewVotes) mapToVotes(voteList VoteList) [][]Vote {
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
	Votes NewVotes `json:"votes"`
}

type NewInputList struct {
	Candidates VoteList   `json:"candidates"`
	Voters     []NewVoter `json:"votes"`
	InverseMap map[int]Vote
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
		Voters:     newVoters,
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
