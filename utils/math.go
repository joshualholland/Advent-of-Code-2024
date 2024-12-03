package utils

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func Occurences(slice []int, value int) int {
	counter := 0

	for _, v := range slice {
		if v == value {
			counter++
		}
	}
	return counter
}
