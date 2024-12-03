package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("real.txt")
	if err != nil {
		panic(err)
	}

	str := string(data)
	numSafe := CountSafe(str)
	fmt.Printf("Number of safe reports %d\n", numSafe)

	offOneSafe := CountOneOffSafe(str)
	fmt.Printf("Number of off one safe reports %d\n", offOneSafe)

}

func mustAtoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Invalid number: %v", err))
	}
	return val
}

func isSafe(levels []string) bool {
	safe := true
	asc := mustAtoi(levels[0]) < mustAtoi(levels[1])
	for i := 0; i < len(levels)-1; i++ {
		if asc {
			dif := mustAtoi(levels[i+1]) - mustAtoi(levels[i])
			if dif < 1 || dif > 3 {
				// fmt.Printf("round %d unsafe, should be asc but got %d diff for pos %d %d\n", round_num, dif, i, i+1)
				safe = false
				break
			}
		} else {
			dif := mustAtoi(levels[i]) - mustAtoi(levels[i+1])
			if dif < 1 || dif > 3 {
				// 	fmt.Printf("round %d unsafe, should be des but got %d diff for pos %d %d\n", round_num, dif, i, i+1)
				safe = false
				break
			}
		}
	}

	return safe
}

func CountSafe(data string) int {
	rounds := strings.Split(string(data), "\n")
	rounds = rounds[:len(rounds)-1]

	count := 0

	for _, round := range rounds {
		levels := strings.Fields(round)
		if isSafe(levels) {
			count++
		}

	}
	return count
}

func CountOneOffSafe(data string) int {
	rounds := strings.Split(string(data), "\n")
	rounds = rounds[:len(rounds)-1]

	count := 0

	for _, round := range rounds {
		levels := strings.Fields(round)
		if isSafe(levels) {
			count++
		} else {
			for i := 0; i < len(levels); i++ {
				newLevel := make([]string, 0, len(levels)-1)
				newLevel = append(newLevel, levels[:i]...)
				newLevel = append(newLevel, levels[i+1:]...)
				if isSafe(newLevel) {
					count++
					break
				}
			}
		}

	}
	return count
}
