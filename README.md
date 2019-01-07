# Markov Chain Text Generator written in Golang
A simple library that takes in text and uses it to create random sentences based on the probability of one or more words leading to the next.

## Usage
Create a mctg instance.  n is the number of words per group.  The higher the number the more like the source text it'll be.
```
myMctg := mctg.New(n)
```
 Load a text file as a corpus and load it into the dictionary.  Can be called multiple times.
```
myMctg.LoadCorpus("input.txt")
```
Generate a paragraph with 5 sentences. Returns a string.
```
myMctg.GenerateParagraph(5)
```
