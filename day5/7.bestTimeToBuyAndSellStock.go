package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
    minP := math.MaxInt32
	maxProfit := 0
	for _ ,p := range prices {
		if p < minP {
			minP = p
		}
		profit := p - minP
		if profit > maxProfit{
			maxProfit = profit
		}
	}
	return maxProfit
}

func main(){
	nums := []int{7,1,5,3,6,4}
	p := maxProfit(nums)
	fmt.Println(p)
}