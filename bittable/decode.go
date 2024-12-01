package bittable

import (
	"fmt"
	"strconv"

	"github.com/ilius/bip39-bittable/bip39"
)

func WordFromBits(bitstr string) string {
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

func WordPairFromBits(bitstr string) (string, string) {
	return WordFromBits(bitstr[:11]), WordFromBits(bitstr[11:])
}
