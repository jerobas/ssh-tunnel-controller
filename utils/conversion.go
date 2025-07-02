package utils

import (
	"log"
	"strconv"
)

func AtoiOrFatal(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Panicf("unexpected non-integer from ps output: %q", str)
	}

	return num
}
