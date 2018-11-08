package main

import (
	"errors"
	"fmt"
	"sort"
)

func calc(inputList InputList) (Result, error) {
	//Check if too few voters
	if len(inputList.Voters) < 1 {
		return Result{}, errors.New("must have at least one vote in voting")
	}

	//Quickly return if just one voter
	if len(inputList.Voters) == 1 {
		return Result{inputList.Voters[0].Votes}, nil
	}

	//Check if any voter has double votes
	for _, voter := range inputList.Voters {
		if voter.hasDoubles() {
			return Result{}, errors.New("voters cannot vote two times at the same item")
		}
	}

	for i := 1; i < len(inputList.Voters); i++ {
		first := inputList.Voters[i-1]
		second := inputList.Voters[i]

		if first.missesVotes(second) {
			return Result{}, errors.New(fmt.Sprintf("someone forgot to vote for all alternatives:\nfirst:  %1v\nsecond: %2v", first, second))
		}
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

func (voter Voter) hasDoubles() bool {
	votes := voter.flattenVotes()
	for i, vote := range votes {
		if vote.existsInList(votes[i+1:]) {
			return true
		}
	}

	return false
}

func (voter Voter) missesVotes(other Voter) bool {
	myVotes := voter.flattenVotes()
	otherVotes := other.flattenVotes()

	if len(myVotes) != len(otherVotes) {
		return true
	}

	for _, myVote := range myVotes {
		if !myVote.existsInList(otherVotes) {
			return true
		}
	}

	return false
}

func (vote Vote) existsInList(votes []Vote) bool {
	for _, v := range votes {
		if v == vote {
			return true
		}
	}
	return false
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
