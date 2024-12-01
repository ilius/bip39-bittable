package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ilius/bip39-bittable/bip39"
)

func wordToBits(word string) string {
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

func main() {
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	words := strings.Split(strings.TrimSpace(string(inputBytes)), " ")
	if len(words)%2 == 1 {
		panic("Must be an even number of words")
	}
	zeros := 0
	ones := 0
	for wordI := 0; wordI < len(words); wordI += 2 {
		line := wordToBits(words[wordI]) + wordToBits(words[wordI+1])
		fmt.Println(line)
		zeros += strings.Count(line, "0")
		ones += strings.Count(line, "1")
	}
	fmt.Printf("%d zeros, %d ones\n", zeros, ones)
}
