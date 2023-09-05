package main

import (
	"strconv"
	"strings"
)

type SpeSkillTest interface {
	NarcissticNumber(string) bool
	ParityOutlier()
	NeedleInAHaystack([]string, string) int
	TheBlueOceanReverse([]int, []int) []int
}

type speSkillTest struct{}

func NewSpeSkillTest() *speSkillTest {
	return &speSkillTest{}
}

func powNumber(num1, num2 int) int {
	result := 1
	for i := 0; i < num2; i++ {
		result *= num1
	}

	return result
}

func (s *speSkillTest) NarcissticNumber(strNumber string) bool {
	eachNumbers := strings.Split(strNumber, "")

	lenNumeric := len(eachNumbers)
	total := 0

	for _, number := range eachNumbers {
		parseNumber, err := strconv.Atoi(number)
		if err != nil {
			return false
		}

		total += powNumber(parseNumber, lenNumeric)
	}

	strTotal := strconv.Itoa(total)

	return strTotal == strNumber
}

func (s *speSkillTest) ParityOutlier() {

}

func (s *speSkillTest) NeedleInAHaystack(haystack []string, needle string) int {
	for i, value := range haystack {
		if strings.EqualFold(value, needle) {
			return i
		}
		// if strings.ToUpper(value) == strings.ToUpper(needle) {
		// 	return i
		// }
	}

	return -1
}

func (s *speSkillTest) TheBlueOceanReverse(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int]bool)

	for _, val := range arr2 {
		arr2Map[val] = true
	}

	var result []int

	for _, val := range arr1 {
		if !arr2Map[val] {
			result = append(result, val)
		}
	}

	return result
}
