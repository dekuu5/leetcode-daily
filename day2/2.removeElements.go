package main

import "fmt"

	func removeElement(nums []int, val int) int {
		i :=0
		for _, v := range nums {
			if v != val {
				nums[i] = v
				i++;
			}
		}
		return i
	}


func main(){
	nums := []int{1,2,2,5,0}
	val := 2;
	k := removeElement(nums, val);
	fmt.Println(nums);
	fmt.Println(k)
}