package bittable

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ilius/bip39-bittable/bip39"
)

func WordToBits(word string) string {
	value, ok := bip39.FindWord(word)
	if !ok {
		panic(fmt.Sprintf("bad word %#v", word))
	}
	bitstr := strconv.FormatUint(uint64(value), 2)
	if len(bitstr) > 11 {
		panic(bitstr)
	}
	bitstr = strings.Repeat("0", 11-len(bitstr)) + bitstr
	return bitstr
}

func WordPairToBits(word1 string, word2 string) string {
	return WordToBits(word1) + WordToBits(word2)
}
