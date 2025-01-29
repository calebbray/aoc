package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type program struct {
	rules rulebook
	pages []pageSet
}

type (
	rulebook map[int][]int
	pageSet  []int
)

func SolveTest(content string) (int, int) {
	return partOne(content), partTwo(content)
}

func partOne(content string) int {
	p, err := newProgram(content)
	if err != nil {
		panic("could not parse program")
	}

	total := 0

	for _, row := range p.ValidSets() {
		total += row.middle()
	}
	return total
}

func partTwo(content string) int {
	p, err := newProgram(content)
	if err != nil {
		panic("could not parse program")
	}

	total := 0
	for _, row := range p.FixedSets() {
		total += row.middle()
	}
	return total
}

func (p program) ValidSets() []pageSet {
	valid := []pageSet{}
	for _, ps := range p.pages {
		if ps.isValid(p.rules) {
			valid = append(valid, ps)
		}
	}
	return valid
}

func (p program) InvalidSets() []pageSet {
	invalid := []pageSet{}
	for _, ps := range p.pages {
		if !ps.isValid(p.rules) {
			invalid = append(invalid, ps)
		}
	}
	return invalid
}

func (p program) FixedSets() []pageSet {
	fixed := []pageSet{}
	for _, ps := range p.pages {
		if !ps.isValid(p.rules) {
			fixed = append(fixed, ps.fixSet(p.rules))
		}
	}
	return fixed
}

func (ps pageSet) middle() int {
	return ps[len(ps)/2]
}

func (ps pageSet) sortedInsert(p int, rb rulebook) {
	if len(ps) == 0 {
		ps[0] = p
	}

	var idx int
	for i, x := range ps {
		curr, ok := rb[x]
		if !ok || slices.Contains(curr, p) {
		} else {
			// what do we do if it isn't in there?
		}
	}
	ps[idx] = p
}

func (ps pageSet) fixSet(rb rulebook) pageSet {
	final := pageSet{}
	for _, p := range ps {
		final.sortedInsert(p, rb)
	}
	return final
	// for i := 0; i < len(ps); i++ {
	// 	curr := ps[j]
	// 	ruleset, ok := rb[curr]
	// 	if !ok {
	// 		continue
	// 	}
	// 	for j := 0; j < len(ps)-1-i; j++ {
	// 		idx := slices.Index(ruleset, ps[j])
	// 		if idx == -1 {
	// 			continue
	// 		}
	//
	// 	}
	// }
	// for i := len(ps) - 1; i >= 0; i-- {
	// 	curr := ps[i]
	// 	check := ps[:i]
	// 	ruleset, ok := rb[curr]
	// 	if !ok {
	// 		continue
	// 	}
	//
	// 	for j := len(check) - 1; j >= 0; j-- {
	// 		idx := slices.Index(ruleset, check[j])
	// 		if idx != -1 {
	// 			return false
	// 		}
	// 	}
	return pageSet{}
}

func (ps pageSet) isValid(rb rulebook) bool {
	for i := len(ps) - 1; i >= 0; i-- {
		curr := ps[i]
		check := ps[:i]
		ruleset, ok := rb[curr]
		if !ok {
			continue
		}

		for j := len(check) - 1; j >= 0; j-- {
			idx := slices.Index(ruleset, check[j])
			if idx != -1 {
				return false
			}
		}
	}

	return true

	//  for i := 0; i < len(rb); i++ {
	//    curr := ps[i]
	//    check := ps[i:]
	//    ruleset, ok := rb[curr]
	//    if !ok {
	//      continue
	//    }
	//
	//    for _, pageCheck := range check {
	//
	//    }
	//  }
	//  idx := slices.Index([])
	// return false
}

func newProgram(content string) (program, error) {
	parts := strings.Split(content, "\n\n")
	if len(parts) != 2 {
		return program{}, fmt.Errorf("invalid content format")
	}

	rb, err := newRulebook(parts[0])
	if err != nil {
		return program{}, err
	}
	pages, err := newPages(parts[1])
	if err != nil {
		return program{}, err
	}
	return program{rules: rb, pages: pages}, nil
}

func newRulebook(rules string) (rulebook, error) {
	rb := make(rulebook)
	rawRules := strings.Split(rules, "\n")
	for i, rule := range rawRules {
		parts := strings.Split(rule, "|")

		before, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid rule format at row %d: %s", i, parts[0])
		}
		after, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid rule format at row %d: %s", i, parts[1])
		}

		r, exists := rb[before]
		if !exists {
			rb[before] = []int{after}
		} else {
			rb[before] = append(r, after)
		}
	}
	return rb, nil
}

func newPages(pages string) ([]pageSet, error) {
	rawPages := strings.Split(pages, "\n")
	p := make([]pageSet, len(rawPages))

	for i, page := range rawPages {
		nums := strings.Split(page, ",")
		pageList := make(pageSet, len(nums))

		for j, stringNumber := range nums {
			num, err := strconv.Atoi(stringNumber)
			if err != nil {
				return nil, fmt.Errorf("invalid page number at row %d: %s", i, stringNumber)
			}

			pageList[j] = num
		}

		p[i] = pageList
	}

	return p, nil
}
