package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	reps := math.Pow(2, 32)

	for num := 0; float64(num) < reps; num++ {
		binary := strconv.FormatInt(int64(num), 2)

		for len(binary)%8 != 0 {
			binary = "0" + binary
		}

		ones := strings.Count(binary, "1")
		zeros := strings.Count(binary, "0")

		if ones == zeros {
			fmt.Println(num, binary)
		}

	}

}
