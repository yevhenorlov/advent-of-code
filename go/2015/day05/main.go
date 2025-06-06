/*

--- Day 5: Doesn't He Have Intern-Elves For This? ---

Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
For example:

ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
jchzalrnumimnmhp is naughty because it has no double letter.
haegwjzuvuyypxyu is naughty because it contains the string xy.
dvszwmarrgswjxmb is naughty because it contains only one vowel.
How many strings are nice?

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	val := []string{}
	for scanner.Scan() {
		val = append(val, scanner.Text())
	}
	return val
}

func containsForbidden(x string) bool {
	forbidden := []string{
		"ab", "cd", "pq", "xy",
	}
	for _, f := range forbidden {
		if strings.Contains(x, f) {
			return true
		}
	}
	return false
}

func containsThreeVowels(x string) bool {
	vowels := "aeiou"
	tally := 0
	for _, v := range x {
		if strings.Contains(vowels, string(v)) {
			tally++
			if tally == 3 {
				return true
			}
		}
	}
	return false
}

// overengineered for more than two :)
func containsTwo(x string) bool {
	if len(x) < 3 {
		return false
	}
	cursor := 0
	lookahead := 1
	tally := 1
	for range x {
		if lookahead == len(x) {
			// out of bounds
			return false
		}
		if x[cursor] == x[lookahead] {
			tally++
		} else {
			// reset
			tally = 1
		}
		if tally == 2 {
			return true
		}
		cursor++
		lookahead++
	}
	return false
}

func getNiceList(input []string) []string {
	niceList := []string{}
	for _, line := range input {
		if containsForbidden(line) {
			continue
		}
		if containsThreeVowels(line) && containsTwo(line) {
			niceList = append(niceList, line)
		}
	}
	return niceList
}

func partOne() {
	input := getInput()
	niceList := getNiceList(input)
	fmt.Println("nice list length is:", len(niceList))
}

/*
PART 2 Task
*/

func partTwo() {
	//
}

func main() {
	partOne()
	partTwo()
}
