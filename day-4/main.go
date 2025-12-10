// https://adventofcode.com/2025/day/4

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
		// fmt.Println(line)

		x = len(strings.Split(line, ""))

		grid = append(grid, strings.Split(line, ""))

		y++
	}

	dir := []struct {
		r int
		c int
	} {
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == "." {
				continue
			}

			tp := 0
			for k := 0; k < 8; k++ {
				if (i + dir[k].r >= 0 &&
					i + dir[k].r < x &&
					j + dir[k].c >= 0 &&
					j + dir[k].c < y) {
					if grid[i + dir[k].r][j + dir[k].c] == "@" {
						tp++
					}
				}
			}

			if tp < 4 {
				res++
			}
		}
	}

	fmt.Println(res)
}

func part2(scanner *bufio.Scanner) {
	prev_res := -1
	res := 0

	grid := [][]string{}
	x := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		x = len(strings.Split(line, ""))

		grid = append(grid, strings.Split(line, ""))

		y++
	}

	dir := []struct {
		r int
		c int
	} {
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}
	
	for prev_res != res {
		prev_res = res
		
		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				if grid[i][j] == "." {
					continue
				}

				tp := 0
				for k := 0; k < 8; k++ {
					if (i + dir[k].r >= 0 &&
						i + dir[k].r < x &&
						j + dir[k].c >= 0 &&
						j + dir[k].c < y) {
						if grid[i + dir[k].r][j + dir[k].c] == "@" {
							tp++
						}
					}
				}

				if tp < 4 {
					grid[i][j] = "."
					res++
				}
			}
		}
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