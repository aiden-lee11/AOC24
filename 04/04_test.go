package main

import (
	"os"
	"testing"
)

func TestPartA(t *testing.T) {
	fileName := "sample.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(data)

	count := countXMAS(str)
	correct_count := 18

	if count != correct_count {
		t.Fatalf("Count expected to be %d but is %d\n", correct_count, count)
	}
}

func TestPartB(t *testing.T) {
	fileName := "sample.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(data)

	count := countX_MAS(str)
	correct_count := 9

	if count != correct_count {
		t.Fatalf("Count expected to be %d but is %d\n", correct_count, count)
	}
}
