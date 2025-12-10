// https://adventofcode.com/2025/day/2

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
//	"regexp"
)

func part1(text string) {
	sum := 0
	// re, _ := regexp.Compile("^(\\d+)\\1{1}$")

	ranges := strings.Split(text, ",")
	for _, r := range ranges {
		start, _ := strconv.Atoi(strings.Split(r, "-")[0])
		end, _ := strconv.Atoi(strings.Split(r, "-")[1])
		length := end - start
		for i := 0; i <= length; i++ {
			num_str := strconv.Itoa(start + i)
			if len(num_str) % 2 == 0 {
				if num_str[0:(len(num_str)/2)] == num_str[(len(num_str)/2):] {
					sum += start + i
				}
			}
			
			// if re.MatchString(num_str) {
			// 	sum += start + i
			// }
		}
	}

	fmt.Println(sum)
}

func part2(text string) {
	sum := 0
	// my regex solution wont work... it would have been so cool...
	// re, _ := regexp.Compile("^(\\d+)\\1+$")

	ranges := strings.Split(text, ",")
	for _, r := range ranges {
		start, _ := strconv.Atoi(strings.Split(r, "-")[0])
		end, _ := strconv.Atoi(strings.Split(r, "-")[1])
		length := end - start
		for i := 0; i <= length; i++ {
			num_str := strconv.Itoa(start + i)
			for j := 2; j <= len(num_str); j++ {
				if len(num_str) % j == 0 {
					interval := len(num_str) / j
					invalid := true
					for k := 1; k < j; k++ {
						if num_str[((k-1)*interval):(k*interval)] != num_str[(k*interval):((k+1)*interval)] {
							invalid = false
						}
					}

					if invalid {
						sum += start + i
						break
					}
				}
			}
			
			// if re.MatchString(num_str) {
			// 	sum += start + i
			// }
		}
	}

	fmt.Println(sum)
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	text := string(file)
	// fmt.Println(text)

	if len(os.Args) > 1 && os.Args[1] == "--part1" {
		part1(text)
	} else if len(os.Args) > 1 && os.Args[1] == "--part2" {
		part2(text)
	} else {
		fmt.Println("invalid argument")
	}
}