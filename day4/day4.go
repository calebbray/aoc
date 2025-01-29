package day4

import (
	"bytes"
	"fmt"
	"strings"

	u "github.com/calebbray/aoc/util"
)

func Solve(content string) (int, int) {
	return partOne(content), partTwo(content)
}

func partOne(content string) int {
	maze := newMaze(content)
	maze.walk()
	return maze.hits
}

func partTwo(content string) int {
	maze := newMaze(content)
	maze.x_mas()
	return maze.hits
}

type xmasMaze struct {
	maze []xmasLine
	hits int
	curr point
}

type xmasLine []byte

type point struct {
	x, y int
}

var directions = []u.Direction{u.Up, u.UpRight, u.Right, u.DownRight, u.Down, u.DownLeft, u.Left, u.UpLeft}

func (m *xmasMaze) scanXmas(last byte, dir u.Direction, currPoint point) {
	nextCoord := point{x: currPoint.x + dir[0], y: currPoint.y + dir[1]}
	if nextCoord.x < 0 || nextCoord.x > len(m.maze[0])-1 || nextCoord.y < 0 || nextCoord.y > len(m.maze)-1 {
		return
	}

	nextChar := m.maze[nextCoord.y][nextCoord.x]
	switch last {
	case 'X':
		if nextChar == 'M' {
			m.scanXmas(nextChar, dir, nextCoord)
		}
	case 'M':
		if nextChar == 'A' {
			m.scanXmas(nextChar, dir, nextCoord)
		}
	case 'A':
		if nextChar == 'S' {
			m.hits++
		}
	}
}

func (m *xmasMaze) next() error {
	if m.curr.x > len(m.maze[0]) && m.curr.y > len(m.maze) {
		return fmt.Errorf("eof")
	}

	// if at end of row increment y and reset x to zero
	if m.curr.x == len(m.maze[0])-1 {
		m.curr.x = 0
		m.curr.y++
		return nil
	}

	// if not at end of row increment x
	if m.curr.x < len(m.maze[0]) {
		m.curr.x++
		return nil
	}

	return nil
}

func (m *xmasMaze) walk() {
	if m.curr.x < 0 || m.curr.x >= len(m.maze[0]) || m.curr.y < 0 || m.curr.y >= len(m.maze) {
		return
	}

	if m.maze[m.curr.y][m.curr.x] == 'X' {
		for _, dir := range directions {
			m.scanXmas('X', dir, m.curr)
		}
	}

	if err := m.next(); err != nil {
		return
	}

	m.walk()
}

var diagonalChecks = []u.Direction{u.UpLeft, u.DownRight, u.UpRight, u.DownLeft}

func (m *xmasMaze) scanMas() {
	checks := make([]byte, 4)
	// Check the four corners around the 'A'
	for i, dir := range diagonalChecks {
		nextY := m.curr.y + dir[1]
		nextX := m.curr.x + dir[0]

		if nextX < 0 || nextX >= len(m.maze[0]) || nextY < 0 || nextY >= len(m.maze) {
			return
		}

		checks[i] = m.maze[nextY][nextX]
	}

	d1, d2 := checks[:2], checks[2:]

	isD1Mas := d1[0] == 'S' && d1[1] == 'M' || d1[0] == 'M' && d1[1] == 'S'
	isD2Mas := d2[0] == 'S' && d2[1] == 'M' || d2[0] == 'M' && d2[1] == 'S'

	if isD1Mas && isD2Mas {
		m.hits++
	}
}

func (m *xmasMaze) x_mas() {
	if m.curr.x < 0 || m.curr.x >= len(m.maze[0]) || m.curr.y < 0 || m.curr.y >= len(m.maze) {
		return
	}

	if m.maze[m.curr.y][m.curr.x] == 'A' {
		m.scanMas()
	}

	if err := m.next(); err != nil {
		return
	}

	m.x_mas()
}

func newMaze(content string) *xmasMaze {
	lines := strings.Split(content, "\n")
	xmas := make([]xmasLine, len(lines))

	for i, line := range lines {
		xmas[i] = xmasLine(line)
	}

	return &xmasMaze{maze: xmas, hits: 0, curr: point{0, 0}}
}

// so I can copy and paste a line
func (x xmasLine) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for i, s := range x {
		out.WriteString(fmt.Sprintf("'%c'", s))
		if i < len(x)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString("},")
	return out.String()
}
