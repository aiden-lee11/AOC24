package main

import (
	"os"
	"testing"
)

func TestSumMultipliers(t *testing.T) {
	fileName := "sample.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(data)

	count := SumMultipliers(str)
	correct_count := 161

	if count != correct_count {
		t.Fatalf("Count expected to be %d but is %d\n", correct_count, count)
	}
}

func TestConditionalSum(t *testing.T) {
	fileName := "sample2.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(data)

	count := ConditionalMultipliers(str)
	correct_count := 48

	if count != correct_count {
		t.Fatalf("Count expected to be %d but is %d\n", correct_count, count)
	}
}
