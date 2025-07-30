package sort

import "fmt"

func findSmallest(sl []int32) int {
	smallest := sl[0]
	smallest_index := 0

	for i := 1; i < len(sl); i++ {
		if sl[i] < smallest {
			smallest = sl[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectionSort(sl1 []int32) {
	fmt.Printf("Original array: %v, len: %d \n", sl1, len(sl1))
	sl1_size := len(sl1)
	sorted_list := []int32{}
	for range sl1_size {
		si := findSmallest(sl1)
		sorted_list = append(sorted_list, sl1[si])
		sl1 = append(sl1[:si], sl1[si+1:]...)
	}
	fmt.Printf("Sorted array: %v, len: %d \n", sorted_list, len(sorted_list))
	fmt.Println()
}

func SelectionSortDemo() {
	sl1 := []int32{10, 20, 11, 18, 16, 15, 9}
	sl2 := []int32{11, 0, 11, -18, 16, 15, 19}
	selectionSort(sl1)
	selectionSort(sl2)
}
