package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ilius/bip39-bittable/bip39"
)

func wordFromBits(bitstr string) string {
	value, err := strconv.ParseUint(bitstr, 2, 16)
	if err != nil {
		panic(err)
	}
	word, ok := bip39.GetWord(int(value))
	if !ok {
		panic(fmt.Sprintf("invalid word index %d with bitstr %s", value, bitstr))
	}
	return word
}

func main() {
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	words := []string{}
	for _, line := range strings.Split(string(inputBytes), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if len(line) != 22 {
			panic(fmt.Sprintf("bad line: %v", line))
		}
		words = append(words, wordFromBits(line[:11]), wordFromBits(line[11:]))
	}
	fmt.Println(strings.Join(words, " "))
}
