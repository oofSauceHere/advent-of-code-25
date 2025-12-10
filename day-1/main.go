// https://adventofcode.com/2025/day/1

package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func part1(scanner *bufio.Scanner) {
	dial := 50
	pwd := 0

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			panic(err)
		}

		if scanner.Text()[0] == 'R' {
			dial = (dial + i) % 100
		} else if scanner.Text()[0] == 'L' {
			dial = (dial - i) % 100
		}

		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			pwd++
		}
	}

	fmt.Println(pwd)
}

func part2(scanner *bufio.Scanner) {
	dial := 50
	pwd := 0

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			panic(err)
		}

		if scanner.Text()[0] == 'R' {
			diff := i + dial
			if dial != 0 {
				diff -= 100
			}
			if diff >= 0 {
				pwd += (diff / 100)
				if dial != 0 {
					pwd++
				}
			}

			dial = (dial + i) % 100
		} else if scanner.Text()[0] == 'L' {
			diff := i - dial
			if diff >= 0 {
				pwd += (diff / 100)
				if dial != 0 {
					pwd++
				}
			}

			dial = (dial - i) % 100
		}

		if dial < 0 {
			dial += 100
		}
	}

	fmt.Println(pwd)
}

func main() {
	fmt.Println("hello chat")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	if len(os.Args) > 1 && os.Args[1] == "--part1" {
		part1(scanner)
	} else if len(os.Args) > 1 && os.Args[1] == "--part2" {
		part2(scanner)
	} else {
		fmt.Println("invalid argument")
	}
}