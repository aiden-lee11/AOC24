package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("real.txt")
	if err != nil {
		panic(err)
	}
	str := string(data)

	closest := CountClosest(str)
	fmt.Println("Closest count:", closest)

	similarity := CountSimilarity(str)
	fmt.Println("Similarity count:", similarity)
}

func CountClosest(data string) int {
	nums := strings.Fields(data)
	nums_length := len(nums)

	left_nums, right_nums := make([]int, nums_length/2), make([]int, nums_length/2)

	for i := 0; i < len(nums)-1; i += 2 {
		left_num, _ := strconv.Atoi(nums[i])
		right_num, _ := strconv.Atoi(nums[i+1])

		left_nums[i/2] = left_num
		right_nums[i/2] = right_num
	}

	sort.Ints(left_nums)
	sort.Ints(right_nums)

	sum_dif := 0
	for i, left := range left_nums {
		sum_dif += int(math.Abs(float64(left - right_nums[i])))
	}

	return sum_dif
}

func CountSimilarity(data string) int {
	nums := strings.Fields(data)
	nums_length := len(nums)

	freqMap := make(map[int]int)

	for i := 1; i < nums_length; i += 2 {
		right_num, _ := strconv.Atoi(nums[i])
		freqMap[right_num]++
	}

	similarity := 0

	for i := 0; i < nums_length; i += 2 {
		left_num, _ := strconv.Atoi(nums[i])
		similarity += left_num * freqMap[left_num]
	}

	return similarity
}
