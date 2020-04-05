package main

import (
	"fmt"
	"os"
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

func getIndex(i, n int) (int, int) {
	return i / n, i % n
}

type NetVote struct {
	candidate  int
	candidate2 int
	diff       int
}

func popular(xs []int, resultMatrix [][]int) func(a, b int) bool {
	return func(a, b int) bool {
		as := 0
		bs := 0
		for _, i := range xs {
			as += resultMatrix[xs[a]][i] - resultMatrix[i][xs[a]]
			bs += resultMatrix[xs[b]][i] - resultMatrix[i][xs[b]]
		}
		return as > bs
	}
}

func less(voteIndexes []int, cycles *UnionFind, resultMatrix [][]int) func(int, int) bool {
	return func(a, b int) bool {
		va := voteIndexes[a]
		vb := voteIndexes[b]
		ra := cycles.Root(va)
		rb := cycles.Root(vb)
		if ra != rb {
			// fmt.Printf("compare %v %v - roots %v %v - res %v %v - sum %v \n",
			// a, b,
			// ra, rb,
			// resultMatrix[ra][rb], resultMatrix[rb][ra],
			// resultMatrix[ra][rb] > resultMatrix[rb][ra],
			// )
			return resultMatrix[ra][rb] > resultMatrix[rb][ra]
		}
		return true
	}
}

func deltaLess(voteIndexes []int, deltaTable []int) func(int, int) bool {
	return func(a, b int) bool {
		va := voteIndexes[a]
		vb := voteIndexes[b]
		// ra := cycles.Root(a)
		// rb := cycles.Root(b)
		// if ra != rb {
		// fmt.Printf("compare %v %v - roots %v %v - res %v %v - sum %v \n",
		// a, b,
		// ra, rb,
		// deltaTable[ra], deltaTable[rb],
		// deltaTable[ra] > deltaTable[rb],
		// )
		return deltaTable[va] > deltaTable[vb]
		// }
		// return true
	}
}

func resolveGroup(xs []int, resultMatrix [][]int) [][]int {
	sort.Slice(xs, popular(xs, resultMatrix))
	gs := groupSorted(xs, popular(xs, resultMatrix))
	var groups [][]int
	for _, g := range gs {
		if len(g) < len(xs) {
			gs = resolveGroup(g, resultMatrix)
			for _, g_ := range gs {
				groups = append(groups, g_)
			}
		} else {
			return gs
		}
	}
	return groups
}

func groupSorted(xs []int, cmp func(int, int) bool) [][]int {
	var grouped [][]int
	var group []int
	group = append(group, xs[0])
	for i := 0; i < len(xs)-1; i++ {
		if cmp(i, i+1) && !cmp(i+1, i) {
			grouped = append(grouped, group)
			group = nil
		}
		group = append(group, xs[i+1])
	}
	if len(group) > 0 {
		grouped = append(grouped, group)
	}
	return grouped
}

func resolve(voteIndexes []int, resultMatrix [][]int) [][]int {
	l := len(voteIndexes)
	preorder := make([][]bool, l)
	for i := range preorder {
		preorder[i] = make([]bool, l)
		preorder[i][i] = true
	}
	updated := true
	for updated {
		updated = false
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				a := resultMatrix[i][j]
				b := resultMatrix[j][i]
				if a <= b {
					if preorder[i][j] == false {
						updated = true
						preorder[i][j] = true
					}
					for k := 0; k < l; k++ {
						if preorder[j][k] == true {
							if preorder[i][k] == false {
								updated = true
								preorder[i][k] = true
							}
						}
					}
				}
				if b <= a {
					if preorder[j][i] == false {
						updated = true
						preorder[j][i] = true
					}
					for k := 0; k < l; k++ {
						if preorder[i][k] == true {
							if preorder[j][k] == false {
								updated = true
								preorder[j][k] = true
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("res matrix: %v\n", resultMatrix)
	fmt.Printf("preorder: %v\n", preorder)

	cycles := New(l)
	for i := range voteIndexes {
		for j := range voteIndexes {
			if preorder[i][j] && preorder[j][i] {
				cycles.Union(i, j)
			}
		}
	}

	fmt.Fprintf(os.Stderr, "cycles: %v\n", *cycles)
	fmt.Fprintf(os.Stderr, "vodeindexes: %v\n", voteIndexes)

	sort.Slice(voteIndexes, less(voteIndexes, cycles, resultMatrix))
	fmt.Fprintf(os.Stderr, "voteindexes after sort: %v\n", voteIndexes)

	// deltaTable := createDeltaTable(voteIndexes, resultMatrix)
	// fmt.Printf("deltas: %v\n", deltaTable)
	// sort.Slice(voteIndexes, deltaLess(voteIndexes, deltaTable))
	// fmt.Fprintf(os.Stderr, "voteindexes after sort2: %v\n", voteIndexes)

	grouped := groupSorted(voteIndexes, less(voteIndexes, cycles, resultMatrix))

	fmt.Fprintf(os.Stderr, "grouped: %v\n", grouped)

	var final [][]int
	for _, g := range grouped {
		gs := resolveGroup(g, resultMatrix)
		for _, g_ := range gs {
			final = append(final, g_)
		}
	}
	fmt.Fprintf(os.Stderr, "%v\n", final)
	return final
}

func createDeltaTable(voteIndexes []int, resultMatrix [][]int) []int {
	l := len(voteIndexes)
	delta := make([]int, l)
	//Wins
	for y := 0; y < l; y++ {
		for x := 0; x < l; x++ {
			delta[x] += resultMatrix[x][y]
		}
	}

	//Losses
	for x := 0; x < l; x++ {
		for y := 0; y < l; y++ {
			delta[y] -= resultMatrix[x][y]
		}
	}
	return delta
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
