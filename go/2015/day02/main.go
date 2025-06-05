/*
--- Day 2: I Was Told There Would Be No Math ---

The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getPaper(sides []int) int {
	return 3*sides[0] + 2*sides[1] + 2*sides[2]
}

func getSides(dimensions []int) []int {
	a := dimensions[0] * dimensions[1]
	b := dimensions[1] * dimensions[2]
	c := dimensions[0] * dimensions[2]
	return []int{a, b, c}
}

func getDimensions(parts []string) []int {
	var dimensions []int
	for _, part := range parts {
		v, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting to int:", err)
		}
		dimensions = append(dimensions, v)
	}
	return dimensions
}

func partOne() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "x")

		dimensions := getDimensions(parts)
		sides := getSides(dimensions)
		slices.Sort(sides) // sort in ascending order
		paper := getPaper(sides)

		fmt.Println("dimensions:", line, "wrap area:", paper)
		total = total + paper
	}
	fmt.Println("total paper needed:", total)
}

/*
 */

func partTwo() {
	//
}

func main() {
	partOne()
	partTwo()
}
