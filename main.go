package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 13, 23, 56, 89, 76}
	find := binarySearch(arr, 23, 0, len(arr)-1)
	fmt.Println(find)

	find = iterBinarySearch(arr, 5, 0, len(arr)-1)
	fmt.Println(find)

}

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {

	if highIndex < lowIndex {
		return -1
	}
	mid := int((highIndex + lowIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex < endIndex {
		mid = int((endIndex + startIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}
