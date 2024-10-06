package main

import "fmt"

func lexicalOrder(n int) []int {
    res:= []int{};
	curr := 1;
	for len(res) < n {
		res = append(res, curr);
		
		if curr * 10 <= n {
			curr *= 10
		}else {
			for curr == n || curr % 10 == 9{
				curr /= 10
			}
			curr +=1
		}
		


	}
	return res;
}

// 122. Best Time to Buy and Sell Stock II
func maxProfit(prices []int) int {
    profint := 0

	for i := 1; i < len(prices); i++ {
		if (prices[i] > prices[i-1]){
			profint+= prices[i] - prices[i-1]
		}
	}
	return profint
}

func main(){
	nums := []int{7,1,5,3,6,4};
	fmt.Println(maxProfit(nums))
}