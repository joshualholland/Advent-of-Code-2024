package solutions

import (
	"advent-code/hello/utils"
	"fmt"
	"log"
	"slices"
)

func Day1Part1() {
	input1, input2, err := utils.GetInputs("./data/lists.txt")

	if err != nil {
		log.Fatal(err)
	}

	diff := 0

	for i := range input1 {
		diff += utils.Abs(input1[i] - input2[i])
	}

	fmt.Println(diff)
}

func Day1Part2() {
	input1, input2, err := utils.GetInputs("./data/lists.txt")

	if err != nil {
		log.Fatal(err)
	}

	similarity := 0

	for i := range input1 {
		if slices.Contains(input2, input1[i]) {
			instances := utils.Occurences(input2, input1[i])
			similarity += input1[i] * instances
		}
	}

	fmt.Println(similarity)
}
