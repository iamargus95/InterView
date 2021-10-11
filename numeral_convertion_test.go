package main

/*

Problem Statement

Write a program to convert a number to Roman number literal and vice versa.
Roman Symbol - Arabic Value
I - 1
V - 5
X - 10
L - 50
C - 100
D - 500
M - 1000

Generally, Roman numerals are written in descending order from left to right,
and are added sequentially, for example MMVI (2006) is interpreted as 1000 + 1000 + 5 + 1.
*/

import (
	"testing"
)

func TestParseInput(t *testing.T) {

	testCases := []struct {
		arabic string
		roman  string
	}{
		{
			arabic: "20",
			roman:  "XX",
		},
		{
			arabic: "50",
			roman:  "L",
		},
		{
			arabic: "101",
			roman:  "CI",
		},
		{
			arabic: "5005",
			roman:  "MMMMMV",
		},
		{
			arabic: "127",
			roman:  "CXXVII",
		},
	}

	for _, tC := range testCases {

		val := arabic2Roman(tC.arabic)
		if tC.roman != val {
			t.Errorf("Expected value is %s, Got %s", tC.roman, val)
		}
	}
}

func TestRoman2Arabic(t *testing.T) {

	testCases := []struct {
		arabic string
		roman  string
	}{
		{
			arabic: "20",
			roman:  "XX",
		},
		{
			arabic: "50",
			roman:  "L",
		},
		{
			arabic: "101",
			roman:  "CI",
		},
		{
			arabic: "5005",
			roman:  "MMMMMV",
		},
		{
			arabic: "127",
			roman:  "CXXVII",
		},
	}

	for _, tC := range testCases {

		val := roman2Arabic(tC.roman)
		if tC.arabic != val {
			t.Errorf("Expected value is %s, Got %s", tC.arabic, val)
		}
	}
}
