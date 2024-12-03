package solutions

import (
	"advent-code/hello/utils"
	"fmt"
	"log"
	"slices"
)

func day1part1() {
	input1, input2, err := utils.GetInputs("./data/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	diff := 0

	for i := range input1 {
		diff += utils.Abs(input1[i] - input2[i])
	}

	fmt.Println("list difference: ", diff)
}

func day1part2() {
	input1, input2, err := utils.GetInputs("./data/day1.txt")

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

	fmt.Println("list similarity: ", similarity)
}

func Day1() {
	day1part1()
	day1part2()
}
