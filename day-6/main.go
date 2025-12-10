// https://adventofcode.com/2025/day/6

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func filter(arr []string) []string {
	res := []string { }

	for _, n := range arr {
		if n != "" {
			res = append(res, n)
		}
	}

	return res
}

func part1(scanner *bufio.Scanner) {
	res := 0

	sheet := [][]string { }
	x := 0
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := filter(strings.Split(line, " "))
		sheet = append(sheet, row)

		x = len(row)
		y++
	}

	for i := 0; i < x; i++ {
		ans := 0
		op := sheet[y-1][i]
		if op == "*" {
			ans = 1
		}

		for j := y-2; j >= 0; j-- {
			n, _ := strconv.Atoi(sheet[j][i])
			if op == "*" {
				ans *= n
			} else {
				ans += n
			}
		}

		res += ans
	}

	fmt.Println(res)
}

func part2(scanner *bufio.Scanner) {
	res := 0

	sheet := [][]string { }
	x := 0
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := filter(strings.Split(line, ""))
		sheet = append(sheet, row)

		x = len(row)
		y++
	}

	op := "*"
	ans := 0
	for i := 0; i < x; i++ {
		if sheet[y-1][i] != " " {
			op = sheet[y-1][i]
			if op == "*" {
				ans = 1
			}
		}

		digit := false
		num := 0
		for j := 0; j < y-1; j++ {
			if sheet[j][i] != " " {
				digit = true
				n, _ := strconv.Atoi(sheet[j][i])
				num *= 10
				num += n
			}
		}

		if digit {
			if op == "*" {
				ans *= num
			} else {
				ans += num
			}
		} else {
			res += ans
			ans = 0
			continue
		}
	}

	fmt.Println(res + ans)
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