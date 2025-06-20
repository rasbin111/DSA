package main

import (
	"fmt"
)

func middle(a []int) int {
	var mid int
	if len(a)%2 == 0 {
		mid = len(a) / 2
	} else {
		mid = (len(a) / 2) + 1
	}
	return mid
}

func binary_search(a []int, searchNum int) {
	low := 0
	high := len(a) - 1

	midIndex := (low + high) / 2
	midValue := a[midIndex]
	found := true

	for searchNum != midValue {
		if searchNum > midValue {
			low = midIndex + 1
		} else {
			high = midIndex - 1
		}

		if low > high {
			found = false
			fmt.Printf("Index of %d is: %d\n", searchNum, -1)
			break
		}

		midIndex = (low + high) / 2
		midValue = a[midIndex]
	}
	if found {

		fmt.Printf("Index of %d is: %d\n", searchNum, midIndex)
	}
}

func main() {
	a := []int{4, 6, 7, 9, 15, 20, 24, 26, 27, 29, 31}
	binary_search(a, 6)
	binary_search(a, 27)
	binary_search(a, 30)

}
