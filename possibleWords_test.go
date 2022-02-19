package main

import (
	"testing"
)

func TestGreens(t *testing.T) {
	testCases := []struct {
		w        string
		g        string
		expected bool
	}{
		{"house", "h....", true},
		{"house", ".o...", true},
		{"house", "..u..", true},
		{"house", "...s.", true},
		{"house", "....e", true},
		{"house", "house", true},
		{"house", "..o...", false},
	}

	for _, testCase := range testCases {
		answer := greens(testCase.w, testCase.g)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s %s expected %t, got %t", testCase.w, testCase.g, testCase.expected, answer)
		}
	}
}

func TestYellows(t *testing.T) {
	testCases := []struct {
		w        string
		y        []string
		expected bool
	}{
		{"house", []string{}, true},
		{"house", []string{"a...."}, false},
		{"house", []string{"a....", "....z"}, false},
		{"house", []string{"a....", "....z", ".a...", "..z.."}, false},
	}

	for _, testCase := range testCases {
		answer := yellows(testCase.w, testCase.y)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s %v expected %t, got %t", testCase.w, testCase.y, testCase.expected, answer)
		}
	}
}

func TestSortUniq(t *testing.T) {
	testCases := []struct {
		s1       string
		s2       string
		s3       []string
		expected string
	}{
		{"house", "dodge", []string{}, "deghosu"},
		{"house", "dodge", []string{"a...."}, "adeghosu"},
		{"house", "dodge", []string{"a....", "....z"}, "adeghosuz"},
		{"house", "dodge", []string{"a....", "....z", ".a...", "..z.."}, "adeghosuz"},
	}

	for _, testCase := range testCases {
		answer := sortUniq(testCase.s1, testCase.s2, testCase.s3)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s %s %v expected %s, got %s", testCase.s1, testCase.s2, testCase.s3, testCase.expected, answer)
		}
	}
}
