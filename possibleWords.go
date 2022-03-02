package main

// go fmt && golint && go test && go run possibleWords.go -green= -grey= -yellow=

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

var (
	green  = flag.String("green", "", "letters known to be correct and in the correct position")
	grey   = flag.String("grey", "", "letters that have not been eliminated and are not green or yellow")
	yellow = flag.String("yellow", "", "letters known to be correct, but in the correct position of the form ..y..,.t....")
)

// greens returns true if the word is not invalidated by the green pattern
func greens(word, green string) bool {
	for i := range green {
		if green[i] == '.' {
			continue
		}
		if word[i] != green[i] {
			return false
		}
	}

	return true
}

// yellows returns true if the word is not invalidated by the yellow patterns
func yellows(word string, yellow []string) bool {
	// word is not allow to have any letters in positions that match a yellow mask
	for _, s := range yellow {
		for i := range s {
			if s[i] == '.' {
				continue
			}
			if word[i] == s[i] {
				return false
			}
		}
	}

	// word must have all of the letters in the yellow masks
	ymap := map[rune]bool{}

	for _, s := range yellow {
		for _, val := range s {
			ymap[val] = true
		}
	}

	delete(ymap, '.')

	for _, val := range word {
		delete(ymap, val)
	}

	return len(ymap) == 0
}

// sortUniq returns a string of the unique letters in s1, s2, s3 in sorted order
func sortUniq(s1, s2 string, s3 []string) string {
	letters := map[rune]bool{}

	for _, val := range s1 {
		letters[val] = true
	}

	for _, val := range s2 {
		letters[val] = true
	}

	for _, s := range s3 {
		for _, val := range s {
			letters[val] = true
		}
	}

	delete(letters, '.')

	l := []rune{}
	for key := range letters {
		l = append(l, key)
	}

	sort.Slice(l, func(i int, j int) bool { return l[i] < l[j] })

	return string(l)
}

func permute(green, grey string, yellow []string) {
	letters := sortUniq(green, grey, yellow)

	for _, a := range letters {
		for _, b := range letters {
			for _, c := range letters {
				for _, d := range letters {
					for _, e := range letters {
						word := fmt.Sprintf("%c%c%c%c%c", a, b, c, d, e)
						if !yellows(word, yellow) {
							continue
						}
						if !greens(word, green) {
							continue
						}
						fmt.Println(word)
					}
				}
			}
		}
	}
}

func main() {
	flag.Parse()

	if *green == "" && *grey == "" && *yellow == "" {
		fmt.Println("You must specify -green, -grey, or -yellow")
		return
	}

	permute(*green, *grey, strings.Split(*yellow, ","))
}
