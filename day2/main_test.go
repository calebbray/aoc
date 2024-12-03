package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplice(t *testing.T) {
	t.Run("removes the third element", func(t *testing.T) {
		s := Level{1, 2, 3, 4, 5}
		want := Level{1, 2, 4, 5}
		got := Splice(s, 2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("unexpected splice behavior: got=(%+v), want=(%+v)", got, want)
		}
	})

	t.Run("iterating and splicing", func(t *testing.T) {
		s := Level{1, 2, 3, 4, 5}
		tests := []Level{
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
				got := Splice(s, i)
				if !reflect.DeepEqual(tests[i], got) {
					t.Errorf("unexpected splice behavior: got=(%+v), want=(%+v)", got, tests[i])
				}
			})
		}
	})
}
