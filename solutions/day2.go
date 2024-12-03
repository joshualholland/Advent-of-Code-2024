package solutions

import (
	"advent-code/hello/utils"
	"fmt"
	"log"
)

func between(lower int, upper int, value int) bool {
	diff := utils.Abs(value)
	return diff > lower && value <= upper
}

func IncreasingOrDecreasing(slice []int) bool {
	var direction int

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != slice[i+1] {
			if slice[i] < slice[i+1] {
				if direction == 0 {
					direction = 1
				} else if direction == -1 {
					return false
				}
			} else if slice[i] > slice[i+1] {
				if direction == 0 {
					direction = -1
				} else if direction == 1 {
					return false
				}
			}

			return true
		}
	}

	return false
}

func day2part1() {
	safetyCounter := 0
	rows, err := utils.ConvertRowsToSlices("./data/day2.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		if IncreasingOrDecreasing(row) {
			for i := range row {
				if between(0, 3, row[i]) {
					safetyCounter++
				}
			}
		}
	}

	fmt.Println(safetyCounter)
}

func Day2() {
	day2part1()
}
