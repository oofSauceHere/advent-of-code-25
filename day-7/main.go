// https://adventofcode.com/2025/day/7

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	res := 0

	grid := [][]string{}
	x := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		x = len(strings.Split(line, ""))

		grid = append(grid, strings.Split(line, ""))

		if y > 0 {
			for i := 0; i < x; i++ {
				if grid[y-1][i] == "S" {
					grid[y][i] = "|"
				} else if grid[y-1][i] == "|" {
					if grid[y][i] == "^" {
						res++
						if i-1 >= 0 {
							grid[y][i-1] = "|"
						}
						if i+1 < x {
							grid[y][i+1] = "|"
						}
					} else {
						grid[y][i] = "|"
					}
				}
			}
		}

		y++
	}

	fmt.Println(res)
}

func part2(scanner *bufio.Scanner) {
	res := 0

	grid := [][]string{}
	x := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		x = len(strings.Split(line, ""))

		grid = append(grid, strings.Split(line, ""))

		cnt := 0
		split := false
		if y > 0 {
			for i := 0; i < x; i++ {
				if grid[y-1][i] == "S" {
					grid[y][i] = "|"
				} else if grid[y-1][i] == "|" {
					if grid[y][i] == "^" {
						split = true
						if i-1 >= 0 {
							if grid[y][i-1] != "|" {
								cnt++
							}
							grid[y][i-1] = "|"
						}
						if i+1 < x {
							if grid[y][i+1] != "|" {
								cnt++
							}
							grid[y][i+1] = "|"
						}
					} else {
						if grid[y][i] != "|" {
							cnt++
						}
						grid[y][i] = "|"
					}
				}
			}
		}

		if split {
			res += cnt
		}

		y++
	}

	fmt.Println(res)
}

func main() {
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