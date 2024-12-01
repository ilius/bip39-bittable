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
	words := strings.Split(strings.TrimSpace(string(inputBytes)), " ")
	if len(words)%2 == 1 {
		panic("Must be an even number of words")
	}
	zeros := 0
	ones := 0
	for wordI := 0; wordI < len(words); wordI += 2 {
		line := bittable.WordPairToBits(words[wordI], words[wordI+1])
		fmt.Println(line)
		zeros += strings.Count(line, "0")
		ones += strings.Count(line, "1")
	}
	fmt.Printf("%d zeros, %d ones\n", zeros, ones)
}
