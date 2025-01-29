package day6

import (
	"fmt"
	"os"
	"strings"

	"github.com/calebbray/aoc/util"
)

type Maze struct {
	maze             util.Maze
	direction        util.Direction
	curr             util.Point
	currObstruction  *util.Point
	distinctCount    int
	obstructionCount int
	seen             map[util.Point]util.Direction
	isStuck          bool
}

func Solve(content string) (int, int) {
	return partOne(content), partTwo(content)
}

func partOne(content string) int {
	m := newMaze(content)
	m.walk(nil)
	return m.distinctCount
}

func partTwo(content string) int {
	tmp := newMaze(content)
	total := 0
	for y, row := range tmp.maze {
		for x := range row {
			m := newMaze(content)
			if m.maze[y][x] == '#' {
				continue
			}
			m.maze[y][x] = '0'
			m.currObstruction = &util.Point{X: x, Y: y}
			m.walkObstructions()
			if m.isStuck {
				total += 1
			}
		}
	}
	// m := newMaze(content)
	// m.currObstruction = &util.Point{X: 7, Y: 9}
	// m.maze[9][7] = '0'
	//
	// seen := make(map[util.Point]util.Direction)
	// m.walk(seen)
	// fmt.Println(m.maze)
	return total
}

func (m *Maze) rotate() {
	switch m.direction {
	case util.Up:
		m.direction = util.Right
	case util.Right:
		m.direction = util.Down
	case util.Down:
		m.direction = util.Left
	case util.Left:
		m.direction = util.Up
	}
}

func (m *Maze) walk(seen map[util.Point]util.Direction) {
	b, done := m.peek()
	if done {
		m.mark('X')
		return
	}

	if b == '#' || b == '0' {
		m.rotate()
	}

	m.mark('X')
	if seen != nil && m.hasSeen(seen) {
		m.isStuck = true
		m.obstructionCount += 1
		return
	}
	m.next()
	m.walk(seen)
}

func (m *Maze) walkObstructions() {
	seen := make(map[util.Point]util.Direction)
	m.walk(seen)
	// done := m.obstruct()
	// if !done {
	// 	m.walkObstructions()
	// }
}

func (m *Maze) obstruct() bool {
	if m.currObstruction == nil {
		m.currObstruction = &util.Point{X: 0, Y: 0}
		m.maze[0][0] = '0'
		return false
	}

	// Reset this guy back to a dot
	m.maze[m.currObstruction.Y][m.currObstruction.X] = '.'
	done := m.nextObstruction(m.currObstruction)
	if done {
		return true
	}
	return false
}

func (m *Maze) nextObstruction(current *util.Point) bool {
	nextPoint := &util.Point{}
	nextX := current.X + 1
	nextY := current.Y
	// If we are at the end of the row, reset X
	// Check if the next row is off the map - if it is we are done
	if nextX >= len(m.maze[0]) {
		nextX = 0
		if nextY+1 >= len(m.maze) {
			return true
		}
		nextY += 1
	}

	nextPoint.X = nextX
	nextPoint.Y = nextY

	if m.maze[nextPoint.Y][nextPoint.X] == '#' {
		m.nextObstruction(nextPoint)
		return false
	}

	m.maze[nextPoint.Y][nextPoint.X] = '0'
	m.currObstruction = nextPoint
	return false
}

func (m *Maze) hasSeen(seen map[util.Point]util.Direction) bool {
	seenDirection, ok := seen[m.curr]
	if !ok {
		seen[m.curr] = m.direction
		return false
	}
	if seenDirection.String() == m.direction.String() {
		return true
	}
	return false
}

func (m *Maze) mark(b byte) {
	currByte := m.maze[m.curr.Y][m.curr.X]
	if currByte != b {
		m.maze[m.curr.Y][m.curr.X] = b
		m.distinctCount++
	}
}

func (m *Maze) peek() (byte, bool) {
	nextY := m.curr.Y + m.direction[1]
	nextX := m.curr.X + m.direction[0]

	if nextX < 0 || nextX >= len(m.maze[0]) || nextY < 0 || nextY >= len(m.maze) {
		return '0', true
	}

	return m.maze[m.curr.Y+m.direction[1]][m.curr.X+m.direction[0]], false
}

func (m *Maze) next() {
	m.curr = util.Point{Y: m.curr.Y + m.direction[1], X: m.curr.X + m.direction[0]}
}

func newMaze(content string) *Maze {
	var start util.Point
	lines := strings.Split(content, "\n")
	maze := make([]util.MazeLine, len(lines))

	for i, line := range lines {
		l := make(util.MazeLine, len(line))
		for j, b := range line {
			if b == '^' {
				start = util.Point{X: j, Y: i}
				l[j] = 'X'
				continue
			}
			l[j] = byte(b)
		}
		maze[i] = l
	}

	return &Maze{curr: start, maze: maze, direction: util.Up, distinctCount: 1, isStuck: false}
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Location [2]int

func ReadFile(path string) []string {
	bytes, _ := os.ReadFile(path)

	return strings.Split(string(bytes), "\n")
}

func getRotatedDirection(direction Direction) Direction {
	switch direction {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}

	panic("Should not have gotten here")
}

func Part2() {
	lines := ReadFile("./day6/input.txt")
	size := len(lines)

	grid := make([][]string, len(lines))
	for row, line := range lines {
		grid[row] = strings.Split(line, "")
	}

	startLoc := Location{0, 0}

	// Find the guard
	for row := range grid {
		for col := range grid[row] {
			c := grid[row][col]
			if c == "^" {
				startLoc = Location{row, col}
				break
			}
		}
	}

	// Array of all available locations to put an obstacle
	openLocations := []Location{}
	for row := range grid {
		for col := range grid[row] {
			loc := grid[row][col]

			if loc == "." {
				openLocations = append(openLocations, Location{row, col})
			}
		}
	}

	total := 0

	for _, openLoc := range openLocations {
		direction := UP
		currLoc := Location{startLoc[0], startLoc[1]}
		nextLoc := Location{currLoc[0], currLoc[1]}

		// Locations the guard has been to already. The key is a string of:
		// `ROW_COL_DIRECTION`. For example, `4_5_2`. Direction is essentially an
		// enum number. The value is just a bool because it doesn't matter, we
		// essentially just want a set.
		pastLocations := make(map[string]bool)

		// Put an obstacle there
		grid[openLoc[0]][openLoc[1]] = "#"

		loopFound := false

		// run the full loop to see if the guard escapes
		for {
			if direction == UP {
				nextLoc = Location{currLoc[0] - 1, currLoc[1]}
			} else if direction == RIGHT {
				nextLoc = Location{currLoc[0], currLoc[1] + 1}
			} else if direction == DOWN {
				nextLoc = Location{currLoc[0] + 1, currLoc[1]}
			} else if direction == LEFT {
				nextLoc = Location{currLoc[0], currLoc[1] - 1}
			}

			// The guard has exited the grid, so we can exit the loop
			if nextLoc[0] < 0 || nextLoc[0] >= size || nextLoc[1] < 0 || nextLoc[1] >= size {
				break
			}

			// We've ran into an obstacle and gotta rotate
			if grid[nextLoc[0]][nextLoc[1]] == "#" {
				direction = getRotatedDirection(direction)
				continue
			}

			key := fmt.Sprintf("%d_%d_%d", currLoc[0], currLoc[1], direction)
			if pastLocations[key] {
				// Break out of the loop because we've been to the same location with
				// the same direction before, which means we're in a loop
				loopFound = true
				break
			}

			pastLocations[key] = true

			currLoc = Location{nextLoc[0], nextLoc[1]}
		}

		if loopFound {
			total += 1
		}

		// take away the obstacle
		grid[openLoc[0]][openLoc[1]] = "."
	}

	fmt.Print("\n\n", "*** total ***", "\n", total, "\n\n\n")
}
