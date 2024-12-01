package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ilius/bip39-bittable/bittable"
)

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
		word1, word2 := bittable.WordPairFromBits(line)
		words = append(words, word1, word2)
	}
	fmt.Println(strings.Join(words, " "))
}
