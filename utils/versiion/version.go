package versiion

import (
	"strconv"
	"strings"
)

func CompareVersion(left string, right string, sep string) int {
	if left == "" && right == "" {
		return 0
	} else if left == "" {
		return -1
	} else if right == "" {
		return 1
	}

	leftArray := strings.Split(left, sep)
	rightArray := strings.Split(right, sep)

	leftLength := len(leftArray)
	rightLength := len(rightArray)
	min := leftLength
	if min > rightLength {
		min = rightLength
	}

	index := 0
	for index < min {
		l, _ := strconv.Atoi(leftArray[index])
		r, _ := strconv.Atoi(rightArray[index])
		if l > r {
			return 1
		} else if l < r {
			return -1
		}
		index++
	}

	if min < leftLength {
		// left is longer
		return 1
	} else if min < rightLength {
		return -1
	}

	return 0
}
