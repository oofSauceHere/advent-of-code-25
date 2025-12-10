// https://adventofcode.com/2025/day/5

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

type id_range struct {
	start, end int
}

func part1(scanner *bufio.Scanner) {
	res := 0

	ranges := []id_range { }

	proc_ids := false
	for scanner.Scan() {
		line := scanner.Text()
		
		if line == "" {
			proc_ids = true
			continue
		}

		if proc_ids {
			n, _ := strconv.Atoi(line)
			for i := 0; i < len(ranges); i++ {
				if n >= ranges[i].start && n <= ranges[i].end {
					res++
					break
				}
			}
		} else {
			start, _ := strconv.Atoi(strings.Split(line, "-")[0])
			end, _ := strconv.Atoi(strings.Split(line, "-")[1])
			ranges = append(ranges, id_range{start, end})
		}
	}

	fmt.Println(res)
}

func part2(scanner *bufio.Scanner) {
	res := 0

	ranges := []id_range { }

	for scanner.Scan() {
		line := scanner.Text()
		
		if line == "" {
			break
		}

		start, _ := strconv.Atoi(strings.Split(line, "-")[0])
		end, _ := strconv.Atoi(strings.Split(line, "-")[1])
		ranges = append(ranges, id_range{start, end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	comb_ranges := []id_range { }

	for i := 0; i < len(ranges); i++ {
		if len(comb_ranges) == 0 || comb_ranges[len(comb_ranges)-1].end < ranges[i].start {
			comb_ranges = append(comb_ranges, ranges[i])
		} else {
			comb_ranges[len(comb_ranges)-1].end = max(ranges[i].end, comb_ranges[len(comb_ranges)-1].end)
		}
	}

	for i := 0; i < len(comb_ranges); i++ {
		res += comb_ranges[i].end - comb_ranges[i].start + 1
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