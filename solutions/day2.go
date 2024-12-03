package solutions

import (
	"advent-code/hello/utils"
	"fmt"
	"log"
)

func day2part1() {
	safetyCounter := 0
	rows, err := utils.ConvertRowsToSlices("./data/day2.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		if (utils.IncreasingSlice(row) || utils.DecreasingSlice(row)) &&
			utils.Max(3, row) &&
			utils.Min(1, row) {
			safetyCounter++
		}
	}

	fmt.Println(safetyCounter)
}

func Day2() {
	day2part1()
}
