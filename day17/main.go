package main

import (
	"fmt"
	"sort"
)



func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	l,r:= 0,len(skill)-1;
	s := -1
	total := 0
	for l <= r{
		sum := skill[l] + skill[r]
		if sum != s && s != -1 {
			return -1
		} 
		s = sum
		total += skill[l] * skill[r]
		l++
		r--
	}
	return int64(total)
}



func main(){
	testCases := [][]int{
		{1, 2, 3, 4},
		{1, 1, 2, 2},
		{1, 2, 3, 5},
		{1, 1, 1, 1},
	}

	for _, tc := range testCases {
		result := dividePlayers(tc)
		fmt.Printf("dividePlayers(%v) = %d\n", tc, result)
	}
}