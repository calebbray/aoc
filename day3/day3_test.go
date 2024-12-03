package day3

import "testing"

const testString = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const testString2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestDay3(t *testing.T) {
	wantOne, wantTwo := 161, 48
	one, two := Solve(testString2)

	if one != wantOne {
		t.Errorf("result for part one incorrect. got=%d, want=%d", one, wantOne)
	}

	if two != wantTwo {
		t.Errorf("result for part two incorrect. got=%d, want=%d", two, wantTwo)
	}
}

func TestMulsRegex(t *testing.T) {
	want := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	got := getMuls(testString)

	if len(want) != len(got) {
		t.Fatalf("got wrong length of string slice. got=%d, want=%d", len(got), len(want))
	}

	for i := range len(got) {
		if got[i] != want[i] {
			t.Errorf("incorrect string from mul regex: got=%q, want=%q", got[i], want[i])
		}
	}
}

func TestParseMultiple(t *testing.T) {
	want := multiple{x: 3, y: 10}
	got := parseMultiple("mul(3,10)")

	if got.x != want.x {
		t.Errorf("values incorrect: got=%d, want=%d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("values incorrect: got=%d, want=%d", got.y, want.y)
	}
}

func TestParseDosAndDonts(t *testing.T) {
	want := []string{"mul(2,4)", "mul(8,5)"}
	got := parseDosAndDonts(testString2)

	if len(want) != len(got) {
		t.Fatalf("got wrong length of string slice. got=%d, want=%d", len(got), len(want))
	}

	for i := range len(got) {
		if got[i] != want[i] {
			t.Errorf("incorrect string from input parsing: got=%q, want=%q", got[i], want[i])
		}
	}
}
