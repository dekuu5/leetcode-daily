package main

import "fmt"


// this is remoce dups | and ||


func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	s := 0
	for _,v := range nums {
		
		if freq, exists := m[v]; !exists || freq <2{
			m[v]+=1 
			nums[s] = v
			s++;
		}
	}
	return s
}

func main(){
	n := []int{0,0,1,1,1,2,2,3,3,4}
	k := removeDuplicates(n);
	for i := 0; i < k; i++ {
		fmt.Println(n[i])
	}
}