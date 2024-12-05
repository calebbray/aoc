package main

import (
	"flag"
	"fmt"

	"github.com/calebbray/aoc/day2"
	"github.com/calebbray/aoc/day3"
	"github.com/calebbray/aoc/day4"
	"github.com/calebbray/aoc/util"
)

func main() {
	day := flag.Int("day", 1, "specify the day to execute")

	flag.Parse()

	solve(*day)
}

func solve(day int) {
	switch day {
	case 1:
		return
	case 2:
		day2.Solve()
	case 3:
		content := util.GetFileContent("day3/input.txt")
		fmt.Println(day3.Solve(content))
	case 4:
		content := util.GetFileContent("day4/input.txt")
		fmt.Println(day4.Solve(content))
	}
}
