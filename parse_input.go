package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	arabic := roman2Arabic(flag.Arg(0))
	fmt.Println(arabic)
	roman := arabic2Roman(flag.Arg(1))
	fmt.Println(roman)
}

func arabic2Roman(input string) string {

	inputInt, _ := strconv.Atoi(input)
	result := ""

	convert := []struct {
		value int
		roman string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, res := range convert {
		for inputInt >= res.value {
			result += res.roman
			inputInt -= res.value
		}
	}

	return result
}

func roman2Arabic(input string) string {

	romanToArabic := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	roman := []byte(input)

	big := 1000
	arabic := 0
	for _, val := range roman {

		mapToChar := romanToArabic[string(val)]
		if big < mapToChar {
			arabic -= 2 * big
		}

		big = mapToChar
		arabic += big
	}

	result := fmt.Sprintf("%d", arabic)
	return result
}
