/*

--- Day 6: Probably a Fire Hazard ---

Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

For example:

turn on 0,0 through 999,999 would turn on (or leave on) every light.
toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.
After following the instructions, how many lights are lit?

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	Verb  string
	FromX int
	FromY int
	ToX   int
	ToY   int
}

func format(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func parseInstruction(v string) instruction {
	re := regexp.MustCompile(`^(.+?)\s+(\d+),(\d+)\s+through\s+(\d+),(\d+)$`)
	parts := re.FindStringSubmatch(v)
	return instruction{
		Verb:  parts[1],
		FromX: format(parts[2]),
		FromY: format(parts[3]),
		ToX:   format(parts[4]),
		ToY:   format(parts[5]),
	}

}

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

type grid struct {
	value [][]bool
}

func (g *grid) Init() {
	gr := make([][]bool, 1000)
	for x, _ := range gr {
		gr[x] = make([]bool, 1000)
		for y, _ := range gr[x] {
			gr[x][y] = false
		}
	}
	g.value = gr
}

func (g *grid) RunInstructions(input []string) {
	for _, line := range input {
		instr := parseInstruction(line)
		switch instr.Verb {
		case "turn on":
			for x := instr.FromX; x <= instr.ToX; x++ {
				for y := instr.FromY; y <= instr.ToY; y++ {
					g.value[x][y] = true
				}
			}
		case "turn off":
			for x := instr.FromX; x <= instr.ToX; x++ {
				for y := instr.FromY; y <= instr.ToY; y++ {
					g.value[x][y] = false
				}
			}
		case "toggle":
			for x := instr.FromX; x <= instr.ToX; x++ {
				for y := instr.FromY; y <= instr.ToY; y++ {
					g.value[x][y] = !g.value[x][y]
				}
			}
		default:
			log.Fatal("unhnadled verb:", instr.Verb)
		}
	}

}

func (g *grid) CalcLights() int {
	count := 0
	for x := range g.value {
		for y := range g.value[x] {
			if g.value[x][y] == true {
				count++
			}
		}
	}
	return count
}

func partOne() {
	input := getInput()
	g := grid{}
	g.Init()
	g.RunInstructions(input)
	lights := g.CalcLights()
	fmt.Println("lights on", lights)
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
