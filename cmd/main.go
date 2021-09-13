package main

import (
	"github.com/arthurDonizetti/dna/internals/domain"
	"fmt"
)

func main() {
	dna := domain.NewDna(mockSequence())

	fmt.Println("is_simian", dna.IsSimian())
}

func mockSequence() []string {
	return []string {
		"CTGAGA",
		"CTATGC",
		"TATTGT",
		"AGAGGG",
		"CCCCTA",
		"TCACTG",
	}
}
