package main

import (
	"fmt"
	"strings"
)

// groupAnagrams
func key(s string) string{
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a'] += 1
	}

	var b strings.Builder

	for _, c := range count{
		b.WriteString(fmt.Sprintf("%d#",c))

	}
	return b.String()
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, w := range strs {
		k := key(w)
		m[k] = append(m[k], w)
	}
	var result [][]string
	for _,v := range m {
		result = append(result, v)

	}
	return result	
}

// summaryRanges
func summaryRanges(nums []int) []string {

	if len(nums) < 1{
		return []string{}
	}
	var result []string
	start := nums[0]
	ptr := 0 
	for i := 0; i < len(nums) && ptr +1 <  len(nums); i++ {
		if nums[ptr]+1 == nums[i+1] {
			ptr ++
		}else {
			if start == nums[ptr] {
				result = append(result, fmt.Sprintf("%d",start))
			}else {
				result = append(result, fmt.Sprintf("%d->%d",start,nums[ptr]))
			}
			start = nums[ptr+1]
			ptr ++
			
		}

	}
	if start == nums[ptr] {
		result = append(result, fmt.Sprintf("%d",start))
	}else {
		result = append(result, fmt.Sprintf("%d->%d",start,nums[ptr]))
	}
	return result
}
	
func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	result := summaryRanges(nums)
	fmt.Println(result)
}