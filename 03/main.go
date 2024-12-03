package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("real.txt")
	if err != nil {
		panic(err)
	}
	str := string(data)

	sum := SumMultipliers(str)
	fmt.Println("Sum:", sum)

	cond_sum := ConditionalMultipliers(str)
	fmt.Println("Conditional Sum:", cond_sum)
}

func SumMultipliers(data string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	count := 0

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		count += num1 * num2
	}

	return count
}

func ConditionalMultipliers(data string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don't\(\))`)
	matches := re.FindAllStringSubmatch(data, -1)

	count := 0

	e := true

	for _, match := range matches {
		if match[0] == "do()" {
			e = true
		} else if match[0] == "don't()" {
			e = false
		} else if e {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			count += num1 * num2
		}

	}
	return count
}
