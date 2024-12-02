package main

import (
	"fmt"
	"log"
	"os"
	"slices"
)

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func getInputs(filename string) ([]int, []int, error) {
	var input1, input2 []int
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return input1, input2, err
	}
	defer file.Close()

	for {
		var left, right int
		_, err := fmt.Fscanf(file, "%d %d\n", &left, &right)
		if err != nil {
			break
		}
		input1 = append(input1, left)
		input2 = append(input2, right)
	}

	return input1, input2, err
}

func main() {
	input1, input2, err := getInputs("./inputs/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	slices.Sort(input1)
	slices.Sort(input2)

	diff := 0

	for i := 0; i < len(input1) && i < len(input2); i++ {
		diff += abs(input1[i] - input2[i])
	}

	// With Range solution
	// for i := range input1 {
	// 	diff += abs(input1[i] - input2[i])
	// }

	fmt.Println(diff)
}
