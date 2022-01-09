package main

import (
	"fmt"
)

var up bool = false
var right bool = false
var down bool = false
var startPositionX int
var startPositionY int
var upAgain bool = true
var grid [6][8]string
var resArray []string
var coordinates [][]int

func main() {
	fmt.Println("Build Grid: ")
	buildGrid(4, 1)
	fmt.Println()
	n := len(grid)

	findTreasure(4, 1, "", "")
	fmt.Println("Possible treasure point: ")
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Print("\n")
	}

	fmt.Println("\nTreasure coordinates: ")
	fmt.Printf("%v", coordinates)
}

func findTreasure(i int, j int, step string, prevStep string) {
	if grid[i][j] == "#" || (step == "down" && prevStep == "up") || (step == "up" && prevStep == "down") {

		return
	}
	findTreasure(i-1, j, "up", step)

	if step == "up" {
		up = true
	}
	findTreasure(i, j+1, "right", step)

	if step == "right" {
		right = true
	}
	findTreasure(i+1, j, "down", step)

	if step == "down" {
		down = true
	}

	if step == "down" && (prevStep == "right" || prevStep == "down") && up && right && down {
		grid[i][j] = "$"
		coordinates = append(coordinates, []int{i, j})
	}
}

func buildGrid(x int, y int) {
	var i, j int
	for i = 0; i < 6; i++ {
		for j = 0; j < 8; j++ {
			if (i == 0 || i == 5 || j == 0 || j == 7) || (i == 2 && j >= 2 && j <= 4) || (i == 3 && (j == 4 || j == 6)) || (i == 4 && j == 2) {
				grid[i][j] = "#"
				fmt.Print("#")
			} else if i == x && j == y {
				grid[i][j] = "X"
				fmt.Print("X")
			} else {
				grid[i][j] = "."
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
