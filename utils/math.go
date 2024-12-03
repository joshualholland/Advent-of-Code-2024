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

func Max(max int, slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if Abs(slice[i]-slice[i-1]) > max {
			return false
		}
	}

	return true
}

func Min(min int, slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if Abs(slice[i]-slice[i-1]) < min {
			return false
		}
	}

	return true
}

func IncreasingSlice(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			return false
		}
	}
	return true
}

func DecreasingSlice(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}
