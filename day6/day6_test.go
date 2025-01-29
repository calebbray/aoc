package day6

import (
	"testing"

	"github.com/calebbray/aoc/util"
)

const testString = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestDay6(t *testing.T) {
	wantOne, wantTwo := 41, 6
	one, two := Solve(testString)

	if one != wantOne {
		t.Errorf("result for part one incorrect. got=%d, want=%d", one, wantOne)
	}

	if two != wantTwo {
		t.Errorf("result for part two incorrect. got=%d, want=%d", two, wantTwo)
	}
}

func TestObstructing(t *testing.T) {
	m := newMaze(testString)
	m.obstruct()
	if m.maze[0][0] != '0' {
		t.Errorf("expected the first spot to be obstructed:\n%s\n", m.maze)
	}

	m.obstruct()
	if m.maze[0][1] != '0' {
		t.Errorf("expected obstruction:\n%s\n", m.maze)
	}

	m.obstruct()
	m.obstruct()
	m.obstruct()

	if m.maze[0][4] != '#' {
		t.Errorf("expected to skip obstruction:\n%s\n", m.maze)
	}
	if m.maze[0][5] != '0' {
		t.Errorf("expected obstruction:\n%s\n", m.maze)
	}

	m.obstruct()
	m.obstruct()
	m.obstruct()
	m.obstruct()
	m.obstruct()

	if m.maze[1][0] != '0' {
		t.Errorf("expected obstruction:\n%s\n", m.maze)
	}
}

func TestGettingLooped(t *testing.T) {
	m := newMaze(testString)
	// m.currObstruction = &util.Point{X: 7, Y: 9}
	m.maze[9][7] = '0'

	seen := make(map[util.Point]util.Direction)
	m.walk(seen)

	if !m.isStuck {
		t.Errorf("expected to get stuck in the maze\nObstruction Count: %d\n%s\n", m.obstructionCount, m.maze)
	}
}
