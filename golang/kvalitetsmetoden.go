package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func calcLegacy(inputList InputList) (Result, error) {
	fmt.Println(strconv.Itoa(len(inputList.Voters[0].flattenVotes())) + "x" + strconv.Itoa(len(inputList.Voters)))

	//Only one voter, let's just return
	if len(inputList.Voters) == 1 {
		return Result{inputList.Voters[0].Votes}, nil
	}

	newList := inputList.convertToNew()
	newResult := newList.compareAllAgainstEachOther()
	oldResult := inputList.compareAllAgainstEachOtherOld()

	hasSameOutput := NewResultCompareObject{
		NewVotes:     newResult,
		CandidateMap: newList.InverseMap,
	}.compare(oldResult)

	if hasSameOutput {
		println("Same output!!! YAY")
	} else {
		println("Oh noes... We did not have the same output!")
	}

	resolved := resolve(newList.Candidates, newResult)
	//mapped := resolved.mapToVotes(newList.Candidates)
	return Result{resolved}, nil
}
func calc(inputList NewInputList) (Result, error) {
	fmt.Println(strconv.Itoa(len(inputList.Candidates)) + "x" + strconv.Itoa(len(inputList.Voters)))

	newResult := inputList.compareAllAgainstEachOther()
	resolved := resolve(inputList.Candidates, newResult)

	return Result{resolved}, nil
}

func (newList NewInputList) compareAllAgainstEachOther() NewVotes {
	var flattenLen = len(newList.Candidates)

	var start = time.Now()
	//victoriesAgainst := make(map[VictoryPair]int)
	victoriesAgainst := make([][]int, flattenLen)
	for i := range victoriesAgainst {
		victoriesAgainst[i] = make([]int, flattenLen)
	}

	//Going through all the voters
	//_ = voterIndex
	var voter NewVoter
	var sameVoteIndex int
	var sameVotes, lowerVotes []int
	var myVote, otherVote int
	for _, voter = range newList.Voters {
		//Going going through the sortings of the votes
		for sameVoteIndex, sameVotes = range voter.Votes {
			//Going through the votes on the same place
			//_ = currentVoteIndex
			for _, myVote = range sameVotes {
				//Comparing this vote against all other votes that are on a lower place than this
				lowerVotes = voter.Votes[sameVoteIndex+1:].flatten()
				//_ = otherVoteIndex
				for _, otherVote = range lowerVotes {
					victoriesAgainst[myVote][otherVote]++
				}
			}
		}
	}

	var elapsed = time.Now().Sub(start)
	fmt.Println("Comparing new list took " + elapsed.String())

	fmt.Println(len(victoriesAgainst))
	return victoriesAgainst
}

type NewResultCompareObject struct {
	NewVotes     NewVotes
	CandidateMap map[int]Vote
}

func (newResult NewResultCompareObject) compare(oldResult map[VictoryPair]int) bool {
	for i, outer := range newResult.NewVotes {
		for j, inner := range outer {
			pair := VictoryPair{newResult.CandidateMap[i], newResult.CandidateMap[j]}
			if oldResult[pair] != inner {
				return false
			}
		}
	}
	return true
}

func (list InputList) compareAllAgainstEachOtherOld() map[VictoryPair]int {
	var start = time.Now()
	victoriesAgainst := make(map[VictoryPair]int)

	//Going through all the voters
	//_ = voterIndex
	var voter Voter
	var sameVoteIndex int
	var sameVotes, lowerVotes []Vote
	var myVote, otherVote Vote
	for _, voter = range list.Voters {
		//Going going through the sortings of the votes
		for sameVoteIndex, sameVotes = range voter.Votes {
			//Going through the votes on the same place
			//_ = currentVoteIndex
			for _, myVote = range sameVotes {
				//Comparing this vote against all other votes that are on a lower place than this
				lowerVotes = voter.Votes[sameVoteIndex+1:].flatten()
				//_ = otherVoteIndex
				for _, otherVote = range lowerVotes {
					pair := VictoryPair{myVote, otherVote} //TODO: better
					victoriesAgainst[pair]++
				}
			}
		}
	}

	var elapsed = time.Now().Sub(start)
	fmt.Println("Comparing old list took " + elapsed.String())

	fmt.Println(len(victoriesAgainst))
	return victoriesAgainst
}

//TODO: Check already resolved to avoid infinite loops
//var alreadyResolved NewVotes

//TODO: Sort after compared
func resolve(votes []Vote, result NewVotes) [][]Vote {
	sortedVoteIndexes := make([]int, len(votes))
	for i := range votes {
		sortedVoteIndexes[i] = i
	}

	sort.Slice(sortedVoteIndexes, func(i, j int) bool {
		thisVote := votes[i]
		otherVote := votes[j]

		println(thisVote.Name + " : " + otherVote.Name)

		return result[i][j] > result[j][i]
	})

	//TODO: Stop right here if we already have tried resolving this exact list of votes
	//alreadyResolved.contains(votes) -> return

	//var needsResolving []ResolveRange

	//Fold votes into two dimensional slice
	var folded [][]Vote
	for i := 1; i < len(votes); i++ {
		if result[i][i-1] == result[i-1][i] {
			//TODO: Make sure we have something in the slice already
			folded[len(folded)-1] = append(folded[len(folded)-1], votes[i])
		} else {
			folded = append(folded, []Vote{votes[i]})
		}
	}

	//Create a two dimensional array, where votes that has the same realVictoriesAgainstGroup as
	//each other gets put in the same place.

	return folded
}

type ResolveRange struct {
	From int
	To   int
}

type VictoryPair struct {
	Me    Vote
	Other Vote
}
