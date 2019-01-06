package mctg

import (
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type MCTG struct {
	dictionary map[string][]string
	startingWords []string
	n int
}

func New(n int) (newMCTG *MCTG) {
	newMCTG = &MCTG{n:n}

	return newMCTG
}


func (this *MCTG) LoadCorpus(path string) {
	this.dictionary = make(map[string][]string)
	this.startingWords = make([]string, 0)
	this.n = 1

	rand.Seed(time.Now().Unix())
	fileBytes, err := ioutil.ReadFile(path)
	re := regexp.MustCompile(`\r?\n`)
	input := re.ReplaceAllString(string(fileBytes), " ")

	if err == nil {
		inputSplit := strings.Split(input, " ")
		nextStarter := false
		for i := 0; i < len(inputSplit); i++ {
			word := inputSplit[i]
			if i != len(inputSplit)-1 {
				for y := i + 1; y < len(inputSplit)-1 && y-i < this.n; y++ {
					if y <= len(inputSplit)-1 {
						word += " " + inputSplit[y]
					}
				}

				value := ""
				if i+this.n <= len(inputSplit)-1 {
					value = inputSplit[i+this.n]
				}
				this.dictionary[word] = append(this.dictionary[word], value)
				if nextStarter {
					this.startingWords = append(this.startingWords, word)
					nextStarter = false
				}
			}
			if (strings.Contains(word, ".")) || (strings.Contains(word, "?")) || (strings.Contains(word, "!")) {
				nextStarter = true
				i += this.n - 1
			}
		}
	}
	/*	for i := 0; i < len(startingWords); i++ {
		fmt.Println(startingWords[i])
	}*/

}

func (this *MCTG)GenerateSentence() string {
	//create a random sentence
	done := false
	sentence := ""
	currentWord := ""
	word := ""
	for !done {
		if currentWord == "" && len(sentence) == 0 {
			currentWord = this.startingWords[rand.Intn(len(this.startingWords))]
			word = currentWord
		} else {
			if len(this.dictionary[currentWord]) == 0 {
				done = true
			} else {
				word = this.dictionary[currentWord][rand.Intn(len(this.dictionary[currentWord]))]
				currentWordSplit := strings.Split(currentWord, " ")
				currentWord = ""
				for i := 1; i < len(currentWordSplit); i++ {
					currentWord += currentWordSplit[i] + " "
				}
				currentWord += word
			}
		}
		punctuation := (strings.Contains(word, ".")) || (strings.Contains(word, "?")) || (strings.Contains(word, "!"))
		if punctuation {
			done = true
		}
		sentence += " " + word
	}

	return sentence
}
