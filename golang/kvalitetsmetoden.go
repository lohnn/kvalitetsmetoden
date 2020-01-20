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
		NewVotes: newResult,
		//CandidateMap: newList.InverseMap,
	}.compare(oldResult)

	if hasSameOutput {
		println("Same output!!! YAY")
	} else {
		println("Oh noes... We did not have the same output!")
	}
	indexes := make([]int, len(newResult))
	for i := range newList.Candidates {
		indexes[i] = i
	}
	// resolved := resolve(indexes, newResult)
	//mapped := resolved.mapToVotes(newList.Candidates)
	return Result{}, nil
}
func calc(inputList NewInputList) (NewResult, error) {
	fmt.Println(strconv.Itoa(len(inputList.Candidates)) + "x" + strconv.Itoa(len(inputList.Votes)))

	newResult := inputList.compareAllAgainstEachOther()
	indexes := make([]int, len(newResult))
	for i := range inputList.Candidates {
		indexes[i] = i
	}
	resolved := resolve(indexes, newResult)
	//TODO: Map back from indexes to vote to store correctly

	return NewResult{Result: resolved, Candidates: inputList.Candidates}, nil
}

func (newList NewInputList) compareAllAgainstEachOther() newVote {
	var flattenLen = len(newList.Candidates)

	var start = time.Now()
	//victoriesAgainst := make(map[VictoryPair]int)
	victoriesAgainst := make([][]int, flattenLen)
	for i := range victoriesAgainst {
		victoriesAgainst[i] = make([]int, flattenLen)
	}

	//Going through all the voters
	//_ = voterIndex
	var voter [][]int // NewVoter
	var sameVoteIndex int
	var sameVotes, lowerVotes []int
	var myVote, otherVote int
	for _, voter = range newList.Votes {
		//Going going through the sortings of the votes
		for sameVoteIndex, sameVotes = range voter {
			//Going through the votes on the same place
			//_ = currentVoteIndex
			for _, myVote = range sameVotes {
				//Comparing this vote against all other votes that are on a lower place than this
				lowerVotes = flatten(voter[sameVoteIndex+1:])
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

// NewResultCompareObject is a helper struct for comparing in an output. Will be removed eventually
type NewResultCompareObject struct {
	NewVotes     newVote
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
var alreadyResolved [][]int

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func checkAlreadyResolved(checkAgainst []int) bool {
	var resolveCheck []int
	for _, resolveCheck = range alreadyResolved {
		if equal(resolveCheck, checkAgainst) {
			return true
		}
	}
	return false
}

func resolve(voteIndexes []int, resultMatrix [][]int) [][]int {
	sortedVoteIndexes := make([]int, len(voteIndexes))
	copy(sortedVoteIndexes, voteIndexes)

	sort.Slice(sortedVoteIndexes, func(i, j int) bool {
		// thisVote := votes[i]
		// otherVote := votes[j]
		println()
		println("Index: " + strconv.Itoa(i) + " : " + strconv.Itoa(j))
		// println(thisVote + " : " + otherVote)
		println(strconv.Itoa(resultMatrix[i][j]) + " : " + strconv.Itoa(resultMatrix[j][i]) + " - " + strconv.FormatBool(resultMatrix[i][j] > resultMatrix[j][i]))

		return resultMatrix[i][j] > resultMatrix[j][i]
	})

	//Fold votes into two dimensional slice
	var folded [][]int
	for i := 0; i < len(sortedVoteIndexes); i++ {
		if i == 0 {
			folded = append(folded, []int{sortedVoteIndexes[i]})
			continue
		}

		lastI := sortedVoteIndexes[i-1]
		thisI := sortedVoteIndexes[i]
		thisVictories := resultMatrix[thisI][lastI]
		otherVictories := resultMatrix[lastI][thisI]
		if thisVictories == otherVictories {
			folded[len(folded)-1] = append(folded[len(folded)-1], thisI)
		} else {
			folded = append(folded, []int{thisI})
		}
	}

	isAlreadyResolved := checkAlreadyResolved(sortedVoteIndexes)
	println("Already checked? " + strconv.FormatBool(isAlreadyResolved))
	if isAlreadyResolved {
		return folded
	}
	alreadyResolved = append(alreadyResolved, sortedVoteIndexes)

	println("\nStarting cleanup")
	var needsResolving []int
	var victories [][]int
	for i := 0; i < len(folded); i++ {
		results := folded[i]
		if len(results) > 1 {
			//Two votes are on the same place
			println("Resolving votes on the same place")

			needsResolving = append(needsResolving, results...)
			resolved := resolve(needsResolving, resultMatrix)
			victories = append(victories, resolved...)
			needsResolving = nil
			// needsResolving.From = i + 1
		} else if !winsAgainstLower(results[0], folded[i+1:], resultMatrix) {
			//The vote has not won over all later votes
			println("Did not win against all lower")
			needsResolving = append(needsResolving, results...)
		} else {
			//TODO: Resolve votes in the gap
			//Let's now try to resolve the votes
			if len(needsResolving) > 0 {
				println("Resolving votes around a gap")
				resolved := resolve(needsResolving, resultMatrix)
				victories = append(victories, resolved...)
				needsResolving = nil
			}
			victories = append(victories, results)
		}
	}

	if len(needsResolving) > 0 {
		println("Resolving votes around a gap")
		resolved := resolve(needsResolving, resultMatrix)
		victories = append(victories, resolved...)
		needsResolving = nil
	}

	//Create a two dimensional array, where votes that has the same realVictoriesAgainstGroup as
	//each other gets put in the same place.

	return victories
}

func winsAgainstLower(myIndex int, others [][]int, resultMatrix [][]int) bool {
	for i := range others {
		for _, otherVoteIndex := range others[i] {
			if resultMatrix[myIndex][otherVoteIndex] <= resultMatrix[otherVoteIndex][myIndex] {
				return false
			}
		}
	}
	return true
}

// ResolveRange is a struct that is used when unwrangling the sorted votes to know what sub-parts needs unwrangling
type ResolveRange struct {
	From int
	To   int
}

func (myRange ResolveRange) needsResolve() bool {
	return myRange.From == myRange.To
}

// VictoryPair is a struct for comparing two votes and who is higher ranked than the other
type VictoryPair struct {
	Me    Vote
	Other Vote
}
