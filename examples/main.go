package main

import (
	"fmt"

	"github.com/0xnu/nwunsch"
)

func main() {
	seq1 := "GATTACA"
	seq2 := "GCATGCU"
	match := 1
	mismatch := -1
	gap := -2

	score, align1, align2 := nwunsch.NeedlemanWunsch(seq1, seq2, match, mismatch, gap)

	fmt.Println("Score:", score)
	fmt.Println("Alignment 1:", align1)
	fmt.Println("Alignment 2:", align2)
}
