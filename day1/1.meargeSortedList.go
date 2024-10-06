package day1

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	newArr := make([]int, 0, n+m)
	
	index2 := 0
	i := 0
	for i < m && index2 < n {
		if nums1[i] < nums2[index2] {
			newArr = append(newArr, nums1[i])
			i++
		} else {
			newArr = append(newArr, nums2[index2])
			index2++
		}
	}
	for i < m {
		newArr = append(newArr, nums1[i])
		i++
	}
	for index2 < n {
		newArr = append(newArr, nums2[index2])
		index2++
	}

	copy(nums1, newArr)
}

func main() {
	n := []int{1, 2, 5, 6, 8, 77, 88, 99, 0, 0, 0, 0, 0, 0, 0, 0} // nums1 must have extra space
	m := []int{3, 4, 7, 9, 10, 20, 40, 50}
	merge(n, 8, m, 8)
	fmt.Println(n) // Output will be the merged sorted array
}
