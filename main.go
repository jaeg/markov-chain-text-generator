package main

import (
	"fmt"

	"github.com/jaeg/markov-chain-text-generator/mctg"
)

func main() {
	myMctg := mctg.New(2)
	myMctg.LoadCorpus("input.txt", true)

	for i := 0; i < 5; i++ {
		fmt.Println(i, " : ", myMctg.GenerateParagraph(5), "\n------")
	}
}
