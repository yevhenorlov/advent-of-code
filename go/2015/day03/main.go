/*

--- Day 3: Perfectly Spherical Houses in a Vacuum ---

Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

> delivers presents to 2 houses: one at the starting location, and one to the east.
^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type SantasList struct {
	Visited map[string]struct{}
}

func (s *SantasList) Visit(coords []int) {
	if s.Visited == nil {
		s.Visited = make(map[string]struct{})
	}
	key := fmt.Sprintf("%v, %v", coords[0], coords[1])
	s.Visited[key] = struct{}{}
}

func partOne() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	coords := []int{0, 0}
	santasList := &SantasList{}
	santasList.Visit(coords) // visit first house before beginning
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		switch char {
		case '<':
			coords[0] = coords[0] - 1
		case '>':
			coords[0] = coords[0] + 1
		case 'v':
			coords[1] = coords[1] - 1
		case '^':
			coords[1] = coords[1] + 1
		default:
			break
		}
		santasList.Visit(coords)
	}

	fmt.Println("total houses visited:", len(santasList.Visited))
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
