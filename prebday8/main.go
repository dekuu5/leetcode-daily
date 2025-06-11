package main


import (
	"fmt"
)

// 3442. Maximum Difference Between Even and Odd Frequency I
func maxDifference(s string) int {
	freq := make([]int, 26)
	for _, v := range s {
		freq[int(v)-int('a')]++
	}
	maxOdd := -1
	minEven := int(^uint(0) >> 1)
	for _, v := range freq {
		if v == 0 {
			continue
		}
		if v %2 == 0  {
			if v < minEven {
				minEven = v
			}
		}else {
			if v > maxOdd {
				maxOdd = v
			}
		}
	}

	
	return maxOdd - minEven
}


func main(){
	fmt.Println(maxDifference("aaaaabbc"))
}
