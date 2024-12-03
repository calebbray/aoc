package day2_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/calebbray/aoc/day2"
)

func TestSplice(t *testing.T) {
	t.Run("removes the third element", func(t *testing.T) {
		s := day2.Level{1, 2, 3, 4, 5}
		want := day2.Level{1, 2, 4, 5}
		got := day2.Splice(s, 2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("unexpected splice behavior: got=(%+v), want=(%+v)", got, want)
		}
	})

	t.Run("iterating and splicing", func(t *testing.T) {
		s := day2.Level{1, 2, 3, 4, 5}
		tests := []day2.Level{
			{2, 3, 4, 5},
			{1, 3, 4, 5},
			{1, 2, 4, 5},
			{1, 2, 3, 5},
			{1, 2, 3, 4},
		}

		for i := range len(s) {
			// ts := make(Level, len(s))
			// copy(ts, s)

			t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
				got := day2.Splice(s, i)
				if !reflect.DeepEqual(tests[i], got) {
					t.Errorf("unexpected splice behavior: got=(%+v), want=(%+v)", got, tests[i])
				}
			})
		}
	})
}
