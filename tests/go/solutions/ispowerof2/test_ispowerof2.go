package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	var args []string
	for i := 1; i < 5; i++ {
		args = append(args, strconv.Itoa(i))
	}

	for i := 0; i < 12; i++ {
		args = append(args, strconv.Itoa(z01.RandIntBetween(1, 2048)))
	}

	args = append(args, "1024")
	args = append(args, "")
	args = append(args, "4096")
	args = append(args, "8388608")

	for _, v := range args {
		z01.ChallengeMain("ispowerof2", strings.Fields(v)...)
	}
	z01.ChallengeMain("ispowerof2", "1", "2")
}