package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type Word struct {
	value        string
	startingWord bool
	punctuation  bool
}

var dictionary = make(map[string][]string)
var startingWords = make([]string, 0)
var n = 2

func main() {
	rand.Seed(time.Now().Unix())
	fileBytes, err := ioutil.ReadFile("input.txt")
	re := regexp.MustCompile(`\r?\n`)
	input := re.ReplaceAllString(string(fileBytes), " ")

	if err == nil {
		inputSplit := strings.Split(input, " ")
		nextStarter := false
		for i := 0; i < len(inputSplit); i++ {
			word := inputSplit[i]
			if i != len(inputSplit)-1 {
				for y := i + 1; y < len(inputSplit)-1 && y-i < n; y++ {
					if y <= len(inputSplit)-1 {
						word += " " + inputSplit[y]
					}
				}

				value := ""
				if i+n <= len(inputSplit)-1 {
					value = inputSplit[i+n]
				}
				dictionary[word] = append(dictionary[word], value)
				if nextStarter {
					startingWords = append(startingWords, word)
					nextStarter = false
				}
			}
			if (strings.Contains(word, ".")) || (strings.Contains(word, "?")) || (strings.Contains(word, "!")) {
				nextStarter = true
				i += n - 1
			}
		}
	}
	/*	for i := 0; i < len(startingWords); i++ {
		fmt.Println(startingWords[i])
	}*/

	for i := 0; i < 50; i++ {
		fmt.Println(i, " : ", generateSentence())
	}
}

func generateSentence() string {
	//create a random sentence
	done := false
	sentence := ""
	currentWord := ""
	word := ""
	for !done {
		if currentWord == "" && len(sentence) == 0 {
			currentWord = startingWords[rand.Intn(len(startingWords))]
			word = currentWord
		} else {
			if len(dictionary[currentWord]) == 0 {
				done = true
			} else {
				word = dictionary[currentWord][rand.Intn(len(dictionary[currentWord]))]
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
