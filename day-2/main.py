# to prove the regex solution works

import re

def part1(text):
    sum = 0
    ranges = text.split(",")

    for r in ranges:
        start = int(r.split("-")[0])
        end = int(r.split("-")[1])

        for i in range(start, end+1):
            res = re.match("^(\\d+)\\1{1}$", str(i))
            if res is not None:
                sum += i
    
    print(sum)

def part2(text):
    sum = 0
    ranges = text.split(",")

    for r in ranges:
        start = int(r.split("-")[0])
        end = int(r.split("-")[1])

        for i in range(start, end+1):
            res = re.match("^(\\d+)\\1+$", str(i))
            if res is not None:
                sum += i
    
    print(sum)

def main():
    with open("input.txt", "r") as f:
        text = f.read()
        part1(text)
        part2(text)

if __name__ == "__main__":
    main()