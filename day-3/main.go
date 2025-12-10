// https://adventofcode.com/2025/day/3

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func max(start int, nums []string) (int, int) {
	index := start
	val := nums[0]
	val_i, _ := strconv.Atoi(val)

	for i, n := range nums {
		n_i, _ := strconv.Atoi(n)
		if n_i > val_i {
			val_i = n_i
			index = start + i
		}
	}

	return index, val_i
}

func part1(scanner *bufio.Scanner) {
	jt := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")

		index, v1 := max(0, line_arr[:(len(line_arr)-1)])
		_, v2 := max(index+1, line_arr[(index+1):])

		jt += v1*10 + v2
	}

	fmt.Println(jt)
}

func part2(scanner *bufio.Scanner) {
	jt := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")

		index := -1
		j := 0
		for i := 0; i < 12; i++ {
			disp := 11 - i
			new_index, v := max(index+1, line_arr[(index+1):(len(line_arr)-disp)])
			index = new_index
			j = j*10 + v
		}

		jt += j
	}

	fmt.Println(jt)
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