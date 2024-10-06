package main 

import "fmt"

// 6.majorityElement
func majorityElement(nums []int) int {
    // O(n) with hashmap 
	h := len(nums) /2
	m := make(map[int]int) // m[number] = freq
	for _, v := range nums {
		m[v] +=1
		if m[v] >h {
			return v;
		}
	}
	return -1
}

func rotate(nums []int, k int)  {
	n := len(nums)
	// k = k%n
	rev := func (s, e int)  {
		for s < e {
			nums[s],nums[e] = nums[e],nums[s]
			s++;
			e--; 
		}
	}
	rev(0, n-1);
	rev(0, k-1);
	rev(k, n-1);

}



func main(){
	nums := []int{2,2,1,1,1,2,2}
	rotate(nums, 2)
	fmt.Println(nums)
}