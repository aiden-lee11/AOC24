package main

import (
	"fmt"
	"os"
	"strings"
)

// Main function to read the file and call the solution
func main() {
	data, err := os.ReadFile("real.txt")
	if err != nil {
		panic(err)
	}
	str := string(data)

	partA := countXMAS(str)
	fmt.Println("Total XMAS count:", partA)

	partB := countX_MAS(str)
	fmt.Println("Total XMAS count:", partB)

}

const word = "XMAS"

// Directions are represented as pairs of (row, column) changes
var all_directions = [][2]int{
	{-1, 0},  // Up
	{1, 0},   // Down
	{0, -1},  // Left
	{0, 1},   // Right
	{-1, -1}, // Top-left diagonal
	{-1, 1},  // Top-right diagonal
	{1, -1},  // Bottom-left diagonal
	{1, 1},   // Bottom-right diagonal
}

func countXMAS(data string) int {
	grid := strings.Split(data, "\n")
	grid = grid[:len(grid)-1]
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Helper function to check if "XMAS" can be found starting from (r, c) in the given direction
	checkDirection := func(r, c, dr, dc int) bool {
		for i := 0; i < len(word); i++ {
			newR := r + dr*i
			newC := c + dc*i
			if newR < 0 || newR >= rows || newC < 0 || newC >= cols || grid[newR][newC] != word[i] {
				return false
			}
		}
		return true
	}

	// Iterate over all cells in the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// If the current cell matches the first letter of "XMAS"
			if grid[r][c] == word[0] {
				// Check all directions
				for _, direction := range all_directions {
					dr, dc := direction[0], direction[1]
					if checkDirection(r, c, dr, dc) {
						count++
					}
				}
			}
		}
	}

	return count
}

func countX_MAS(data string) int {
	grid := strings.Split(data, "\n")
	grid = grid[:len(grid)-1]
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Helper function to check if "XMAS" can be found starting from (r, c) in the given direction
	checkDiagnols := func(r, c int) bool {
		valid := func(corner1, corner2 byte) bool {
			return (corner1 == word[1] && corner2 == word[3]) || (corner1 == word[3] && corner2 == word[1])
		}

		// Check top left bottom right then top right bottom left
		return valid(grid[r-1][c-1], grid[r+1][c+1]) && valid(grid[r+1][c-1], grid[r-1][c+1])
	}

	// Iterate over all cells in the grid
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			// If the current cell matches the first letter of "XMAS"
			if string(grid[r][c]) == "A" {
				// Check all directions
				if checkDiagnols(r, c) {
					count++
				}
			}
		}
	}

	return count
}
