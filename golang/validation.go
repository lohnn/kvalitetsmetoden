package main

import (
	"errors"
	"fmt"
)

func (list InputList) validate() error {
	//Check if too few voters
	if len(list.Voters) < 1 {
		return errors.New("must have at least one vote in voting")
	}
	fmt.Println("At least all voters has voted something")

	//Quickly return if just one voter
	if len(list.Voters) == 1 {
		return nil
	}
	fmt.Println("More than one voter... Nice")

	//Check if any voter has double votes
	for _, voter := range list.Voters {
		if voter.hasDoubles() {
			return errors.New("voters cannot vote two times at the same item")
		}
	}
	fmt.Println("Well, all voters has made sure to not have duplicates")

	for i := 1; i < len(list.Voters); i++ {
		first := list.Voters[i-1]
		second := list.Voters[i]

		if first.missesVotes(second) {
			return fmt.Errorf("someone forgot to vote for all alternatives:\nfirst:  %1v\nsecond: %2v", first, second)
		}
	}
	fmt.Println("Everyone voted for everything.")
	return nil
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
	var v Vote
	for _, v = range votes {
		if v == vote {
			return true
		}
	}
	return false
}
