package main

import (
	"fmt"
)

func lexicalOrder(n int) []int {
	res := make([]int, n)
	i := 0 

	var dfa func(int)

	dfa = func(j int) {
		if j > n {
			return 
		}
		res[i] = j
		i++
		j *= 10

		for k := 0; k < 10; k++ {
			dfa(j + k)
		}			
	}

	for k := 1; k < 10; k++ {
		dfa(k)
	}

	return res
}
// stack based 
func lexicalOrder2(n int) []int {
	res := make([]int, n)
	i := 0 
	curr := 1
	for i < n{
		res[i] = curr
		i++

		if curr * 10 <= n{
			curr *= 10
		} else {
			for (curr % 10) == 9 || curr == n {
				curr /= 10

			} 
			curr++
		}
		
	}

			return res
}


func main (){
	fmt.Println("wqwqw")
}