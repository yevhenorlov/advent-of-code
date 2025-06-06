/*

--- Day 4: The Ideal Stocking Stuffer ---

Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

For example:

If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....

*/

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()

}

func getAnswerForPrefix(prefix string) (int, error) {
	input := getInput()
	for i := 0; i < 99999999; i++ {
		val := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(val))
		str := hex.EncodeToString(hash[:])

		if strings.HasPrefix(str, prefix) {
			return i, nil
		}
	}
	return -1, errors.New("limit reached, no value found")

}

func partOne() {
	answer, err := getAnswerForPrefix("00000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("part 1 answer:", answer)
}

/*

--- Part Two ---

Now find one that starts with six zeroes.

*/

func partTwo() {
	answer, err := getAnswerForPrefix("000000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("part 2 answer:", answer)
}

func main() {
	partOne()
	partTwo()
}
