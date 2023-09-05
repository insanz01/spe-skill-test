package main

import "testing"

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

	isNarsis := app.NarcissticNumber("153")

	if !isNarsis {
		t.Errorf("Must expected true, but got %v", isNarsis)
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

	isNarsis := app.NarcissticNumber("153")

	if !isNarsis {
		t.Errorf("Must expected true, but got %v", isNarsis)
	}
}
