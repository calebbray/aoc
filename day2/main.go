package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	Ascending direction = iota
	Descending
)

type Level []int

func NewLevel(str string) Level {
	s := strings.Split(str, " ")
	l := make(Level, len(s))
	for i, c := range s {
		num, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		l[i] = num
	}
	return l
}

func (l Level) isSafeWithoutOne() bool {
	if l.isSafe() {
		return true
	}

	for i := range l {
		localLevel := Splice(l, i)

		if localLevel.isSafe() {
			return true
		}
	}
	return false
}

func (l Level) isSafe() bool {
	return l.isSorted() && l.isGradual()
}

func (l Level) isGradual() bool {
	if !l.isValidLevel() {
		return false
	}

	for i := range len(l) - 1 {
		diff := math.Abs(float64(l[i]) - float64(l[i+1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func (l Level) isSorted() bool {
	if !l.isValidLevel() {
		return false
	}

	d := l.deriveDirection(0, 1)

	for i := range len(l) - 1 {
		currDirection := l.deriveDirection(i, i+1)
		if d != currDirection {
			return false
		}
	}
	return true
}

func (l Level) isValidLevel() bool {
	return len(l) > 1
}

func (l Level) deriveDirection(a, b int) direction {
	if l[a] < l[b] {
		return Ascending
	} else {
		return Descending
	}
}

func Splice(s Level, i int) Level {
	if i > len(s) {
		panic("index out of bounds")
	}

	ts := make(Level, len(s))
	copy(ts, s)

	return append(ts[:i], ts[i+1:]...)
}
func main() {
	content := getFileContent("input.txt")
	strs := strings.Split(content, "\n")

	levels := make([]Level, len(strs))
	for i, str := range strs {
		levels[i] = NewLevel(str)
	}

	safe := 0
	safeWithoutOne := 0

	for _, s := range levels {
		if s.isSafe() {
			safe++
		}
		if s.isSafeWithoutOne() {
			safeWithoutOne++
		}
	}

	fmt.Println(safe)
	fmt.Println(safeWithoutOne)
}

func getFileContent(source string) string {
	b, err := os.ReadFile(source)
	if err != nil {
		panic("an error occurred reading the file")
	}
	return string(b[:len(b)-1])
}
