package main

import (
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	var left []byte
	var right []byte

	b, err := os.ReadFile("test.txt")
	if err != nil {
		panic("an error occurred reading the file")
	}

	for chunk := range slices.Chunk(b, 6) {
		left = append(left, chunk[0])
		right = append(right, chunk[4])
	}

	slices.Sort(left)
	slices.Sort(right)

	var sum float64
	for i := range left {
		sum += math.Abs(float64(left[i]) - float64(right[i]))
	}
	fmt.Println(sum)
}
