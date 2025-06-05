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

type Santa struct {
	coords  [2]int
	Visited map[string]struct{}
}

func (s *Santa) Walk(direction rune) {
	switch direction {
	case '<':
		s.coords[0] = s.coords[0] - 1
	case '>':
		s.coords[0] = s.coords[0] + 1
	case 'v':
		s.coords[1] = s.coords[1] - 1
	case '^':
		s.coords[1] = s.coords[1] + 1
	default:
		break
	}
}

func (s *Santa) Visit() {
	if s.Visited == nil {
		s.Visited = make(map[string]struct{})
	}
	key := fmt.Sprintf("%v, %v", s.coords[0], s.coords[1])
	s.Visited[key] = struct{}{}
}

func partOne() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	santa := Santa{
		coords: [2]int{0, 0},
	}
	santa.Visit() // visit first house before beginning
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		santa.Walk(char)
		santa.Visit()
	}

	fmt.Println("total houses visited:", len(santa.Visited))
}

/*

--- Part Two ---

The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.

*/

func countUniqueKeys(set1, set2 map[string]struct{}) int {
	unique := make(map[string]struct{})

	for key := range set1 {
		unique[key] = struct{}{}
	}
	for key := range set2 {
		unique[key] = struct{}{}
	}
	return len(unique)
}

func partTwo() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	santa := Santa{
		coords: [2]int{0, 0},
	}
	roboSanta := Santa{
		coords: [2]int{0, 0},
	}

	santa.Visit()     // visit first house before beginning
	roboSanta.Visit() // visit first house before beginning
	index := 0
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if index%2 == 0 {
			santa.Walk(char)
			santa.Visit()
		} else {
			roboSanta.Walk(char)
			roboSanta.Visit()
		}
		index++
	}
	fmt.Println("total houses visited by Santa:", len(santa.Visited))
	fmt.Println("total houses visited by RoboSanta:", len(roboSanta.Visited))
	fmt.Println("unique houses visited:", countUniqueKeys(santa.Visited, roboSanta.Visited))
}

func main() {
	partOne()
	partTwo()
}
