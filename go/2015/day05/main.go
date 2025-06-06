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

func getNiceListOne(input []string) []string {
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
	niceList := getNiceListOne(input)
	fmt.Println("part 1 nice list length is:", len(niceList))
}

/*

--- Part Two ---

Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
For example:

qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.
How many strings are nice under these new rules?

*/

func containsTwoPairs(x string) bool {
	if len(x) < 4 {
		// out of bounds
		return false
	}
	for i := range x {
		if i+4 > len(x) {
			// out of bounds
			return false
		}
		pair := x[i : i+2]
		for j := i; j+4 <= len(x); j++ {
			if x[j+2:j+4] == pair {
				return true
			}
		}
	}
	return false
}

func containsASandwich(x string) bool {
	if len(x) < 3 {
		return false
	}
	for i := range x {
		if i+2 >= len(x) {
			return false
		}
		if x[i] == x[i+2] {
			return true
		}
	}
	return false
}

func getNiceListTwo(input []string) []string {
	niceList := []string{}
	for _, line := range input {
		if containsTwoPairs(line) && containsASandwich(line) {
			niceList = append(niceList, line)
		}
	}
	return niceList
}

func partTwo() {
	input := getInput()
	niceList := getNiceListTwo(input)
	fmt.Println("part 2 nice list length is:", len(niceList))
}

func main() {
	partOne()
	partTwo()
}
