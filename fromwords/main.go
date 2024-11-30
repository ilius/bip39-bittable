package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ilius/bip39-coder/bip39"
)

func main() {
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(inputBytes), " ")
	if len(words)%2 == 1 {
		panic("Must be an even number of words")
	}
	zeros := 0
	ones := 0
	for wordI := 0; wordI < len(words); wordI += 2 {
		data := bip39.Decode(words[wordI] + " " + words[wordI+1])
		parts := make([]string, len(data))
		for i, n := range data {
			parts[i] = fmt.Sprintf("%08b", n)
		}
		line := strings.Join(parts, "")
		if line[:2] != "00" {
			panic(line)
		}
		line = line[2:]
		fmt.Println(line)
		zeros += strings.Count(line, "0")
		ones += strings.Count(line, "1")
	}
	fmt.Printf("%d zeros, %d ones\n", zeros, ones)
}
