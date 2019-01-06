package main

import (
	"fmt"
	"sort"
	"strconv"
)

func calc(inputList InputList) (Result, error) {
	fmt.Println(strconv.Itoa(len(inputList.Voters[0].flattenVotes())) + "x" + strconv.Itoa(len(inputList.Voters)))

	e := inputList.validate()
	if e != nil {
		return Result{}, e
	}

	//Only one voter, let's just return
	if len(inputList.Voters) == 1 {
		return Result{inputList.Voters[0].Votes}, nil
	}

	compareAllAgainstEachOther(inputList)

	resolved := resolve(inputList.Voters[0].flattenVotes())
	return Result{resolved}, nil
}

//TODO: Sort after compared
func resolve(votes []Vote) [][]Vote {
	sort.Slice(votes, func(i, j int) bool {
		thisVote := votes[i]
		otherVote := votes[j]

		return thisVote.victoriesAgainstGroup(votes) > otherVote.victoriesAgainstGroup(votes)
	})

	//Create a two dimensional array, where votes that has the same realVictoriesAgainstGroup as
	//each other gets put in the same place.



	return [][]Vote{{votes[0]}}
}

func compareAllAgainstEachOther(list InputList) map[VictoryPair]int {
	victoriesAgainst := make(map[VictoryPair]int)

	//Going through all the voters
	//_ = voterIndex
	for _, voter := range list.Voters {
		//Going going through the sortings of the votes
		for sameVoteIndex, sameVotes := range voter.Votes {
			//Going through the votes on the same place
			//_ = currentVoteIndex
			for _, myVote := range sameVotes {
				//Comparing this vote against all other votes that are on a lower place than this
				lowerVotes := flattenVotes(voter.Votes[sameVoteIndex+1:])
				//_ = otherVoteIndex
				for _, otherVote := range lowerVotes {
					pair := VictoryPair{myVote, otherVote}
					victoriesAgainst[pair]++
				}
			}
		}
	}
	fmt.Println(victoriesAgainst)
	fmt.Println(len(victoriesAgainst))
	return victoriesAgainst
}

func (vote Vote) victoriesAgainstGroup(votes []Vote) int {
	return 1
}

type VictoryPair struct {
	Me    Vote
	Other Vote
}

func (voter Voter) flattenVotes() []Vote {
	return flattenVotes(voter.Votes)
}

func flattenVotes(votesList [][]Vote) []Vote {
	var votes []Vote
	for _, voteArray := range votesList {
		votes = append(votes, voteArray...)
	}
	return votes
}
