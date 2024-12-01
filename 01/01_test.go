package main

import (
	"os"
	"testing"
)

func TestClosestCount(t *testing.T) {
	fileName := "sample.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(data)

	count := CountClosest(str)
	correct_count := 11

	if count != correct_count {
		t.Fatalf("Count expected to be %d but is %d\n", correct_count, count)
	}
}

func TestSimilarityCount(t *testing.T) {
	fileName := "sample.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(data)

	count := CountSimilarity(str)
	correct_count := 31

	if count != correct_count {
		t.Fatalf("Count expected to be %d but is %d\n", correct_count, count)
	}
}
