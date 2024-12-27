package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ilius/bip39-bittable/bittable"
)

func readBits(input io.Reader) []uint32 {
	bits := []uint32{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			return bits
		}
		line = strings.ReplaceAll(line, " ", "")
		if len(line) != 22 {
			panic(fmt.Sprintf("bad line: %v", line))
		}
		bits = append(bits, bittable.Uint32FromBits(line))
	}
	return bits
}

func xorList(a []uint32, b []uint32) []uint32 {
	c := []uint32{}
	bn := len(b)
	bi := 0
	for _, x := range a {
		c = append(c, x^b[bi])
		bi = (bi + 1) % bn
	}
	return c
}

func main() {
	a := readBits(os.Stdin)
	if len(a) == 0 {
		panic("No input")
	}
	b := readBits(os.Stdin)
	if len(b) == 0 {
		panic("Second input block is missing")
	}
	// fmt.Println(a)
	// fmt.Println(b)
	c := xorList(a, b)
	for _, x := range c {
		fmt.Println(bittable.Uint32ToBits(x))
	}
}
