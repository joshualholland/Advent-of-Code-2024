package solutions

import (
	"advent-code/hello/utils"
	"fmt"
	"log"
	"slices"
)

func Day1() {
	input1, input2, err := utils.GetInputs("./data/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	slices.Sort(input1)
	slices.Sort(input2)

	diff := 0

	for i := 0; i < len(input1) && i < len(input2); i++ {
		diff += utils.Abs(input1[i] - input2[i])
	}

	// With Range solution
	// for i := range input1 {
	// 	diff += abs(input1[i] - input2[i])
	// }

	fmt.Println(diff)
}
