package solutions

import (
	"advent-code/hello/utils"
	"fmt"
	"log"
)

func validate(slice []int) bool {
	decreasing := utils.DecreasingSlice(slice)
	increasing := utils.IncreasingSlice(slice)
	max := utils.Max(3, slice)
	min := utils.Min(1, slice)
	return (decreasing || increasing) && max && min
}

func removeLevelValidation(slice []int) bool {
	for i := 0; i < len(slice); i++ {
		slice1 := append([]int{}, slice[:i]...)
		slice1 = append(slice1, slice[i+1:]...)
		if validate(slice1) {
			return true
		}
	}
	return false
}

func Day2() {
	safetyCounter := 0
	rows, err := utils.ConvertRowsToSlices("./data/day2.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		if validate(row) || removeLevelValidation(row) {
			safetyCounter++
		}
	}

	fmt.Println(safetyCounter)
}
