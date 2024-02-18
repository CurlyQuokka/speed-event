package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	v := os.Args[1]
	n, err := strconv.Atoi(v)
	if err != nil {
		log.Panic(err)
	}

	participants := generateParticipants(n)
	generateRounds(participants)
}

func generateParticipants(n int) *[]int {
	participants := []int{}
	for i := 1; i <= n/2+n%2; i++ {
		participants = append(participants, i)
	}

	for i := n; i > n/2+n%2; i-- {
		participants = append(participants, i)
	}

	return &participants
}

func prepareParticipants(p *[]int) int {
	n := len(*p)
	rounds := n - 1
	if n%2 != 0 {
		*p = append(*p, 0)
		rounds = n
	}

	return rounds
}

type pair struct {
	p1, p2 int
}

func (p *pair) isEqual(v *pair) bool {
	return p.p1 == v.p1 && p.p2 == v.p2
}

func generateRounds(participants *[]int) *[]pair {
	rounds := prepareParticipants(participants)
	n := len(*participants)

	half := (n / 2)

	pairs := &[]pair{}

	for i := 0; i < rounds; i++ {
		fmt.Printf("Round: %d \n", i+1)
		for j := 0; j < half; j++ {
			fmt.Printf("%d - %d\n", (*participants)[j], (*participants)[j+half])
			*pairs = append(*pairs, pair{p1: (*participants)[j], p2: (*participants)[j+half]})
		}
		if n == 2 {
			break
		}
		newParticipants := []int{}
		newParticipants = append(newParticipants, (*participants)[0])
		newParticipants = append(newParticipants, (*participants)[half])
		newParticipants = append(newParticipants, (*participants)[1:half-1]...)
		newParticipants = append(newParticipants, (*participants)[half+1:]...)
		newParticipants = append(newParticipants, (*participants)[half-1])
		participants = &newParticipants

		fmt.Println()
	}

	return pairs
}
