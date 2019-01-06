package main

import (
	"fmt"

	"github.com/jaeg/markov-chain-text-generator/mctg"
)

func main() {
	myMctg := mctg.New(2)
	myMctg.LoadCorpus("input.txt")

	for i := 0; i < 50; i++ {
		fmt.Println(i, " : ", myMctg.GenerateParagraph(5))
	}
}
