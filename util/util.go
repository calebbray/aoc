package util

import (
	"bytes"
	"fmt"
	"os"
)

func GetFileContent(source string) string {
	b, err := os.ReadFile(source)
	if err != nil {
		panic("an error occurred reading the file")
	}
	return string(b[:len(b)-1])
}

type Point struct {
	X, Y int
}

type Direction [2]int

var (
	UpLeft    Direction = [2]int{-1, -1}
	Up        Direction = [2]int{0, -1}
	UpRight   Direction = [2]int{1, -1}
	Left      Direction = [2]int{-1, 0}
	Right     Direction = [2]int{1, 0}
	DownLeft  Direction = [2]int{-1, 1}
	Down      Direction = [2]int{0, 1}
	DownRight Direction = [2]int{1, 1}
)

// for debugging
func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case UpRight:
		return "UpRight"
	case Right:
		return "Right"
	case DownRight:
		return "DownRight"
	case Down:
		return "Down"
	case DownLeft:
		return "DownLeft"
	case Left:
		return "Left"
	case UpLeft:
		return "UpLeft"
	default:
		return "invalid direction"
	}
}

type (
	Maze     []MazeLine
	MazeLine []byte
)

// so I can copy and paste a line
func (l MazeLine) String() string {
	var out bytes.Buffer
	out.WriteString("|")
	for _, s := range l {
		out.WriteString(fmt.Sprintf(" %c ", s))
	}
	out.WriteString("|")
	return out.String()
}

func (m Maze) String() string {
	var out bytes.Buffer
	for range m {
		out.WriteString("---")
	}
	out.WriteString("--")
	out.WriteString("\n")
	for _, l := range m {
		out.WriteString(l.String())
		out.WriteString("\n")
	}
	for range m {
		out.WriteString("---")
	}
	out.WriteString("--")
	return out.String()
}
