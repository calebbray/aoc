package day5

import (
	"testing"
)

const testString = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestDay5(t *testing.T) {
	wantOne, wantTwo := 143, 0
	one, two := SolveTest(testString)

	if one != wantOne {
		t.Errorf("result for part one incorrect. got=%d, want=%d", one, wantOne)
	}

	if two != wantTwo {
		t.Errorf("result for part two incorrect. got=%d, want=%d", two, wantTwo)
	}
}

func TestRowValidity(t *testing.T) {
	testRulebook := rulebook{
		29: {13},
		47: {53, 13, 61, 29},
		53: {29, 13},
		61: {13, 53, 29},
		75: {29, 53, 47, 61, 13},
		97: {13, 61, 47, 29, 53, 75},
	}
	t.Run("testing a valid row", func(t *testing.T) {
		testRow := pageSet{75, 47, 61, 53, 29}
		got := testRow.isValid(testRulebook)
		if got == false {
			t.Error("row was not valid.")
		}
	})
	t.Run("testing an invalid row", func(t *testing.T) {
		testRow := pageSet{61, 13, 29}
		got := testRow.isValid(testRulebook)
		if got {
			t.Error("row should not have been valid")
		}
	})
	t.Run("test fixing a row", func(t *testing.T) {
		invalidRow := pageSet{61, 13, 29}
		got := invalidRow.fixSet(testRulebook)
		want := pageSet{61, 29, 13}
		t.Logf("got=%+v, want=%+v", got, want)
		assertPageset(t, []pageSet{got}, []pageSet{want})
	})

	t.Run("testing a mix of rows", func(t *testing.T) {
		want := []pageSet{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
		}

		p, err := newProgram(testString)
		if err != nil {
			t.Fatalf("could not parse program: %s", err)
		}

		got := p.ValidSets()
		assertPageset(t, got, want)
	})
}

func TestParsingProgram(t *testing.T) {
	want := program{
		rules: rulebook{
			29: {13},
			47: {53, 13, 61, 29},
			53: {29, 13},
			61: {13, 53, 29},
			75: {29, 53, 47, 61, 13},
			97: {13, 61, 47, 29, 53, 75},
		},
		pages: []pageSet{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		},
	}

	got, err := newProgram(testString)
	if err != nil {
		t.Fatalf("could not parse program: %s", err)
	}

	assertValidRuleset(t, got.rules, want.rules)
	assertPageset(t, got.pages, want.pages)
}

func assertPageset(t testing.TB, got, want []pageSet) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("pages created not length of wanted. got=%d, want=%d", len(got), len(want))
	}

	for i, pages := range want {
		if len(want[i]) != len(got[i]) {
			t.Fatalf("pagesets are not the same length at set %d. got=%d, want=%d", i, len(got[i]), len(want[i]))
		}
		for j := range pages {
			t.Log()
			if want[i][j] != got[i][j] {
				t.Errorf("page number is not the same in %d page set: got=%d, want=%d", i+1, want[i][j], got[i][j])
			}
		}
	}
}

func assertValidRuleset(t testing.TB, got, want rulebook) {
	t.Helper()
	for key, value := range want {
		exists, ok := got[key]
		if !ok {
			t.Errorf("missing rule set: %d not in created rulebook", key)
		}

		for i := range value {
			if len(value) != len(exists) {
				t.Errorf("length of rules not the same for %d key. got=%d, want=%d", key, len(exists), len(want))
			}

			if value[i] != exists[i] {
				t.Errorf("Expected to get rule for %d key, got=%d, want=%d", key, exists[i], value[i])
			}
		}
	}
}
