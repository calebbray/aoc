package day3

import (
	"fmt"
	"regexp"
)

type multiple struct {
	x, y int
}

type lexer struct {
	input     string
	position  int
	recording bool
}

func newLexer(input string) *lexer {
	return &lexer{input: input, recording: true}
}

func (l *lexer) parseToggle() {
	startPos := l.position
	if l.recording {
		chunk := l.input[startPos : startPos+7]
		if chunk == "don't()" {
			l.recording = false
			l.position += 7
			return
		}
	} else {
		chunk := l.input[startPos : startPos+4]
		if chunk == "do()" {
			l.position += 4
			l.recording = true
			return
		}
	}
	l.position++
}

func (l *lexer) parseMul() (string, bool) {
	startPos := l.position
	var mul []byte
	if l.input[startPos:startPos+4] != "mul(" {
		l.position++
		return "", false
	}
	l.position += 4
	startPos += 4
	mul = append(mul, 'm', 'u', 'l', '(')

	for isDigit(l.input[l.position]) {
		mul = append(mul, l.input[l.position])
		l.position++
	}

	if l.input[l.position] != ',' {
		l.position++
		return "", false
	}

	l.position++
	mul = append(mul, ',')

	for isDigit(l.input[l.position]) {
		mul = append(mul, l.input[l.position])
		l.position++
	}

	if l.input[l.position] != ')' {
		l.position++
		return "", false
	}

	l.position++
	mul = append(mul, ')')

	return string(mul), true
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *lexer) consume() []string {
	var mulstrs []string
	for l.position < len(l.input) {
		ch := l.input[l.position]
		switch ch {
		case 'd':
			l.parseToggle()
		case 'm':
			if l.recording {
				mul, ok := l.parseMul()
				if ok {
					mulstrs = append(mulstrs, mul)
				}
			} else {
				l.position++
			}
		default:
			l.position++
		}
	}
	return mulstrs
}

func Solve(content string) (int, int) {
	return partOne(content), partTwo(content)
}

func partOne(content string) int {
	total := 0
	multiples := parseMultiples(getMuls(content))

	for _, m := range multiples {
		total += m.factor()
	}
	return total
}

func partTwo(content string) int {
	total := 0
	multiples := parseMultiples(parseDosAndDonts(content))

	for _, m := range multiples {
		total += m.factor()
	}
	return total
}

func (m multiple) factor() int {
	return m.x * m.y
}

func getMuls(content string) []string {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return r.FindAllString(content, -1)
}

func parseDosAndDonts(content string) []string {
	l := newLexer(content)
	return l.consume()
}

func parseMultiples(muls []string) []multiple {
	mults := make([]multiple, len(muls))
	for i, mul := range muls {
		mults[i] = parseMultiple(mul)
	}
	return mults
}

func parseMultiple(mul string) multiple {
	var x, y int
	_, err := fmt.Sscanf(mul, "mul(%d,%d)", &x, &y)
	if err != nil {
		fmt.Println(err)
		panic("issue scanning tuple")
	}
	return multiple{x, y}
}
