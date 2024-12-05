package day4

import (
	"fmt"
	"testing"
)

const testString = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestDay4(t *testing.T) {
	wantOne, wantTwo := 18, 9
	one, two := Solve(testString)

	if one != wantOne {
		t.Errorf("result for part one incorrect. got=%d, want=%d", one, wantOne)
	}

	if two != wantTwo {
		t.Errorf("result for part two incorrect. got=%d, want=%d", two, wantTwo)
	}
}

const simpleTest = `XMAS
MMXX
AXAX
SXXS`

const backwardTest = `SXXS
XAXA
XXMM
SAMX`

const downAndLeft = `XXXX
XXMX
XAXX
SXXX`

const upAndRight = `XXXS
XXAX
XMXX
XXXX`

const sam = `SXS
XAX
MXM`

func TestMazeRunning(t *testing.T) {
	t.Run("testing next() functionality", func(t *testing.T) {
		maze := newMaze(simpleTest)

		assertPoint(t, maze.curr, point{0, 0})

		maze.next()
		assertPoint(t, maze.curr, point{1, 0})

		maze.next()
		maze.next()
		maze.next()
		assertPoint(t, maze.curr, point{0, 1})
	})

	t.Run("test scanning for xmas", func(t *testing.T) {
		tests := []struct {
			name       string
			direction  direction
			testString string
			start      point
		}{
			{"right", Right, simpleTest, point{0, 0}},
			{"down", Down, simpleTest, point{0, 0}},
			{"down and right", DownRight, simpleTest, point{0, 0}},
			{"left", Left, backwardTest, point{3, 3}},
			{"up", Up, backwardTest, point{3, 3}},
			{"up and left", UpLeft, backwardTest, point{3, 3}},
			{"down and left", DownLeft, downAndLeft, point{3, 0}},
			{"up and right", UpRight, upAndRight, point{0, 3}},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("scanning to the %s", tt.name), func(t *testing.T) {
				m := newMaze(tt.testString)
				m.scanXmas('X', tt.direction, tt.start)
				if m.hits != 1 {
					t.Errorf("expected to get an xmas hit to the %s", tt.name)
				}
			})
		}
	})

	t.Run("walk the maze and check hits", func(t *testing.T) {
		maze := newMaze(simpleTest)
		maze.walk()
		want := 3
		if maze.hits != want {
			t.Errorf("result incorrect. got=%d, want=%d", maze.hits, want)
		}
	})
	t.Run("walk the maze and check hits backwards", func(t *testing.T) {
		maze := newMaze(backwardTest)
		maze.walk()
		want := 3
		if maze.hits != want {
			t.Errorf("result incorrect. got=%d, want=%d", maze.hits, want)
		}
	})

	t.Run("scanning simple sam diagonal", func(t *testing.T) {
		maze := newMaze(sam)
		maze.x_mas()
		want := 1
		if maze.hits != want {
			t.Errorf("result incorrect. got=%d, want=%d", maze.hits, want)
		}
	})
}

func TestArrayConversion(t *testing.T) {
	want := []xmasLine{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}
	got := newMaze(testString).maze

	if len(want) != len(got) {
		t.Fatalf("got wrong length of xmas map. got=%d, want=%d", len(got), len(want))
	}
}

func assertPoint(t testing.TB, want, got point) {
	t.Helper()
	if want.x != got.x && want.y != got.y {
		t.Errorf("maze is not in the right state: want=(%+v), got=(%+v)", want, got)
	}
}
