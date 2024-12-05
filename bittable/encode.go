package bittable

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ilius/bip39-bittable/bip39"
)

func splitStringIntoChunks(data string, size int) []string {
	var chunk string
	chunks := make([]string, 0, len(data)/size+1)
	for len(data) >= size {
		chunk, data = data[:size], data[size:]
		chunks = append(chunks, chunk)
	}
	if len(data) > 0 {
		chunks = append(chunks, data)
	}
	return chunks
}

func addSpacesToPlainTextLine(line string, size int) string {
	return strings.Join(splitStringIntoChunks(line, size), " ")
}

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
	bitstr = addSpacesToPlainTextLine(bitstr, 4)
	return bitstr
}

func WordPairToBits(word1 string, word2 string) string {
	return WordToBits(word1) + " " + WordToBits(word2)
}
