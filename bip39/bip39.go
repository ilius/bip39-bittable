package bip39

import (
	"bytes"
	"math/big"

	"github.com/ilius/bip39-bittable/bip39/assets"
)

var (
	wordCount = 2048
	bigRadix  = big.NewInt(int64(wordCount))
	bigZero   = big.NewInt(0)
)

var words, indexByWord = loadWords("english.txt")

func init() {
	if len(words) != wordCount {
		panic("Mismatch number of words")
	}
}

func loadWords(fpath string) ([]string, map[string]int16) {
	textB, err := assets.FS.ReadFile("english.txt")
	if err != nil {
		panic(err)
	}
	indexByWord := map[string]int16{}
	words := []string{}
	for wordIndex, wordB := range bytes.Split(textB, []byte("\n")) {
		word := string(wordB)
		if word == "" {
			continue
		}
		wordIndex16 := int16(wordIndex)
		if int(wordIndex16) != wordIndex {
			panic("Too many words")
		}
		indexByWord[word] = wordIndex16
		words = append(words, word)
	}
	return words, indexByWord
}

func WordCount() int {
	return wordCount
}

func GetWord(index int) (string, bool) {
	if index >= wordCount {
		return "", false
	}
	return words[index], true
}

func FindWord(word string) (int, bool) {
	index, ok := indexByWord[word]
	return int(index), ok
}
