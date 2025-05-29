package main

import "sort"


func firstMissingPositive(nums []int) int {
    // O(n^2)
	// maxInt := int(1<<31 - 1)
	// for i := 1; i < maxInt; i++ {
	// 	lowest := true
	// 	for j := 0; j < len(nums); j++ {
	// 		if nums[j] == i {
	// 			lowest = false
	// 		}
	// 	}
	// 	if lowest {
	// 		return i
	// 	}
	// }
	// return 0

	//O(nlog(n))
	// sort.Ints(nums)
	// // l = 1
	// // add as long as it is keep going above l 
	// // [0,2,2,1,1]
	// l := 1
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] < l {
	// 		continue;
	// 	}
	// 	if nums[i] == l {
	// 		l++
	// 	}else if l > nums[i]{
	// 		return l
	// 	}
	// }
	// return l
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num > 0 && num <= n {
			if nums[num-1] > 0 {
				nums[num-1] = -nums[num-1]
			}
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

}
	func abs(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

func main() {
	
}
