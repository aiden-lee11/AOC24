package main

import (
	"os"
	"testing"
)

func TestCountSafe(t *testing.T) {
	fileName := "sample.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(data)

	count := CountSafe(str)
	correct_count := 2

	if count != correct_count {
		t.Fatalf("Count expected to be %d but is %d\n", correct_count, count)
	}
}

func TestCountOneOffSafe(t *testing.T) {
	fileName := "sample.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(data)

	count := CountOneOffSafe(str)
	correct_count := 4

	if count != correct_count {
		t.Fatalf("Count expected to be %d but is %d\n", correct_count, count)
	}
}
