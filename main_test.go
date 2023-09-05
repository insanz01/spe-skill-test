package main

import (
	"testing"
)

func TestNarsissticNumber(t *testing.T) {
	testScenario := []struct {
		Name     string
		Number   string
		Expected bool
		KindTest bool // negative case -> false, positive case -> true
	}{
		{
			Name:     "case 1",
			Number:   "153",
			Expected: true,
			KindTest: true,
		},
		{
			Name:     "case 2",
			Number:   "111",
			Expected: false,
			KindTest: true,
		},
		{
			Name:     "case 3",
			Number:   "121",
			Expected: true,
			KindTest: false,
		},
	}

	app := NewSpeSkillTest()

	for _, tt := range testScenario {
		t.Run(tt.Name, func(t *testing.T) {
			got := app.NarcissticNumber(tt.Number)

			if got != tt.Expected {
				if tt.KindTest {
					t.Errorf("Expected=%v, but Got=%v", tt.Expected, got)
				}
			}
		})
	}
}

func TestNeedleInHaystackNumber(t *testing.T) {
	testScenario2 := []struct {
		Name     string
		Haystack []string
		Search   string
		Expected int
		KindTest bool // negative case -> false, positive case -> true
	}{
		{
			Name:     "case 1",
			Haystack: []string{"red", "blue", "yellow", "black", "grey"},
			Search:   "blue",
			Expected: 1,
			KindTest: true,
		},
		{
			Name:     "case 2",
			Haystack: []string{"red", "blue", "yellow", "black", "grey"},
			Search:   "yellow",
			Expected: 2,
			KindTest: true,
		},
		{
			Name:     "case 3",
			Haystack: []string{"red", "blue", "yellow", "black", "grey"},
			Search:   "white",
			Expected: -1,
			KindTest: true,
		},
		{
			Name:     "case 4",
			Haystack: []string{"red", "blue", "yellow", "black", "grey"},
			Search:   "red",
			Expected: -1,
			KindTest: false,
		},
	}

	app := NewSpeSkillTest()

	for _, tt := range testScenario2 {
		t.Run(tt.Name, func(t *testing.T) {
			got := app.NeedleInAHaystack(tt.Haystack, tt.Search)

			if got != tt.Expected {
				if tt.KindTest {
					t.Errorf("Expected=%d, but Got=%d", tt.Expected, got)
				}
			}
		})
	}
}

func TestBlueOceanReverse(t *testing.T) {
	testScenario2 := []struct {
		Name     string
		Arr1     []int
		Arr2     []int
		Expected []int
		KindTest bool // negative case -> false, positive case -> true
	}{
		{
			Name:     "case 1",
			Arr1:     []int{1, 2, 3, 4, 6, 10},
			Arr2:     []int{1},
			Expected: []int{2, 3, 4, 6, 10},
			KindTest: true,
		},
		{
			Name:     "case 2",
			Arr1:     []int{1, 5, 5, 5, 5, 3},
			Arr2:     []int{5},
			Expected: []int{1, 3},
			KindTest: true,
		},
		{
			Name:     "case 3",
			Arr1:     []int{1, 5, 5, 5, 5, 3},
			Arr2:     []int{0},
			Expected: []int{1, 5, 5, 5, 5, 3},
			KindTest: true,
		},
	}

	app := NewSpeSkillTest()

	for _, tt := range testScenario2 {
		t.Run(tt.Name, func(t *testing.T) {
			got := app.TheBlueOceanReverse(tt.Arr1, tt.Arr2)

			if len(got) != len(tt.Expected) {
				if tt.KindTest == false {
					t.Errorf("length result is not same")
					return
				}
			}

			for i, g := range got {
				if g != tt.Expected[i] {
					if tt.KindTest != true {
						t.Errorf("Expected=%d, but Got=%d", tt.Expected[i], g)
					}
				}
			}
		})
	}
}
