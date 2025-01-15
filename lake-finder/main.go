package main

import (
	"errors"
	"fmt"
)

func main() {

	// First input example
	// Input data: grid dimensions and the given grid values
	m := 4            // m is the number of rows in the grid
	n := 5            // n is the number of columns in the grid
	grid := [][]byte{ // Defining the grid containing map data
		[]byte("#.###"), // First row, '#' for land, '.' for water
		[]byte("..###"), // Second row
		[]byte("##.##"), // Third row
		[]byte("####."), // Fourth row
	}

	/*
		// Second input example
		// Input data: grid dimensions and the given grid values
		m := 4            // m is the number of rows in the grid
		n := 4            // n is the number of columns in the grid
		grid := [][]byte{ // Defining the grid containing map data
			[]byte("####"), // First row, '#' for land, '.' for water
			[]byte("####"), // Second row
			[]byte("####"), // Third row
			[]byte("####"), // Fourth row
		}
	*/

	// Input validation: Check if the grid matches the given dimensions
	if err := validateInput(m, n, grid); err != nil {
		// If there is an input error, display the error message and stop
		fmt.Println("Invalid input:", err)
		return
	}

	// Display the grid dimensions and content
	fmt.Println("Input grid dimensions:", m, n)
	// Display each row of the grid
	for i := 0; i < m; i++ {
		// Convert each row of the grid into a string and display it
		fmt.Println("Grid row:", string(grid[i]))
	}

	// Count and display the number of lakes
	result := countLakes(grid, m, n)
	// Display the result of the lake count
	fmt.Println("Number of lakes:", result)
}

// Function to count the number of lakes
func countLakes(grid [][]byte, m, n int) int {
	// Directions for DFS search (vertical, horizontal, and diagonal)
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1},
	}
	var dfs func(x, y int) // Declare DFS function
	lakeCount := 0         // Variable to count the number of lakes

	// DFS function to explore connected water cells
	dfs = func(x, y int) {
		// Check if the position x, y is within the grid bounds and if it's water (.)
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != '.' {
			return
		}
		// Mark this position as land (#) after visiting
		grid[x][y] = '#'
		// Continue searching in the neighbor directions (vertical, horizontal, diagonal)
		for _, dir := range directions {
			dfs(x+dir[0], y+dir[1]) // Perform DFS to the neighbor direction
		}
	}

	// Loop to search every cell in the grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// If water ('.') is found, we have found a new lake
			if grid[i][j] == '.' {
				lakeCount++ // Increment lake count
				dfs(i, j)   // Explore the lake using DFS
			}
		}
	}

	return lakeCount // Return the number of lakes found
}

// Function to validate input
func validateInput(m, n int, grid [][]byte) error {
	// Ensure the number of rows matches m
	if len(grid) != m {
		// If the number of rows doesn't match, return an error
		return errors.New("the number of rows does not match m")
	}

	// Ensure each row has length n
	for i, row := range grid {
		// If the row length doesn't match n, return an error
		if len(row) != n {
			return fmt.Errorf("row %d does not have the correct length of n", i+1)
		}
	}

	// If all validations pass, return nil (no error)
	return nil
}
