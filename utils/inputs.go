package utils

import (
	"fmt"
	"os"
)

func GetInputs(filename string) ([]int, []int, error) {
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
