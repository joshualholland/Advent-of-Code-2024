package utils

import (
	"fmt"
	"os"
	"slices"
	"strings"
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

	slices.Sort(input1)
	slices.Sort(input2)

	return input1, input2, err
}

func ConvertRowsToSlices(filename string) ([][]int, error) {
	var slices [][]int
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return slices, err
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	for _, line := range lines {
		var row []int
		fields := strings.Fields(line)
		for _, field := range fields {
			var num int
			_, err := fmt.Sscanf(field, "%d", &num)
			if err != nil {
				return nil, fmt.Errorf("failed to parse number: %v", err)
			}
			row = append(row, num)
		}
		slices = append(slices, row)
	}

	return slices, err
}
