package main

import (
	"fmt"
	"math"
	"sort"
)

// type Stack struct {
// 	arr []string
// }

// func (s *Stack) Push(v string) {
// 	s.arr = append(s.arr, v)
// }
// func (s *Stack) Pop() (string,bool) {
// 	if len(s.arr) == 0 {
// 		return "",false
// 	}
// 	v := s.arr[len(s.arr)-1]
// 	s.arr = s.arr[:len(s.arr)-1]
// 	return v , true
// }

// func (s *Stack)isEmpty() bool {
// 	return len(s.arr) ==0
// }

// func simplifyPath(path string) string {
//     s := Stack{}
// 	arr := strings.Split(path, "/")
// 	for _, v := range arr {
// 		if v == "" || v == "."{
// 			continue
// 		}else if v == ".." {
// 			if !s.isEmpty() {
// 				s.Pop()
// 			}
// 		}else {
// 			s.Push(v)
// 		}

// 	}
// 	return "/" + strings.Join(s.arr, "/")

// }
// 121. Best Time to Buy and Sell Stock

func maxProfit(prices []int) int {
	maxP := 0

	mStock := math.MaxInt


	for i := 0; i < len(prices); i++ {
		if prices[i] < mStock {
			mStock = prices[i]
		}else {
			maxP = max(maxP , prices[i] - mStock)
		}
	}
	return maxP
}
// 20. Valid Parentheses

type Stack struct {
	arr []rune
}

func (s *Stack) Push(v rune) {
	s.arr = append(s.arr, v)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}
	v := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return v, true
}

func (s *Stack) Peek() (rune, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}
	return s.arr[len(s.arr)-1], true
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}
func isValid(s string) bool {
	ss := make([]rune, len(s))
	top:=0
	for _, v := range s {
		if v == '(' ||  v == '{' ||  v == '[' {
			ss[ top ] = v
			top++
		} else {
			if top == 0 {
				return false
			}
			c := ss[top-1]
			if  (v == ')' && c == '(') ||
				(v == '}' && c == '{') ||
				(v == ']' && c == '[') {
					top--
			}else {
				return false
			}
		}
	}
	return top == 0
	
}
// 209. Minimum Size Subarray Sum
// brute force
// func minSubArrayLen(target int, nums []int) int {
// 	n := len(nums)
// 	minLen := n + 1
// 	for w := 1; w <= n; w++ {
// 		for i := 0; i <= n-w; i++ {
// 			sum := 0
// 			for j := i; j < i+w; j++ {
// 				sum += nums[j]
// 			}
// 			if sum >= target {
// 				if w < minLen {
// 					minLen = w
// 				}
// 				break
// 			}
// 		}
// 		if minLen <= w {
// 			break
// 		}
// 	}

// 	if minLen == n+1 {
// 		return 0
// 	}
// 	return minLen
// }
// O(n)
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLen := n+1
	sum := 0
	l := 0

	for r := 0; r < n; r++ {
		sum += nums[r];

		for sum >= target {
			if r-l+1 < minLen {
				minLen = r -l +1 
			}
			sum -= nums[l]
			l++
		}
	}
	if minLen == n+1 {
		return 0
	}
	return minLen
}

// 15. 3Sum
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int,0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		l,r := i+1, len(nums)-1
		
		for l < r{
			sum := nums[l] + nums[r] + nums[i]
			switch {
			case sum < 0:
				l++
			case sum > 0 :
				r--
			default:
			
				result = append(result, []int{nums[i], nums[l], nums[r]})
				
				for	l < r && nums[l] == nums[l-1] {
					l ++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				l++
				r--
			}

		} 
		
	}
	return result
}


func main(){
	// Test simplifyPath function
	// fmt.Println("Testing simplifyPath:")
	// fmt.Println(simplifyPath("/home/")) // Should print: /home
	// fmt.Println(simplifyPath("/../")) // Should print: /
	// fmt.Println(simplifyPath("/home//foo/")) // Should print: /home/foo
	// fmt.Println(simplifyPath("/a/./b/../../c/")) // Should print: /c

	// Test maxProfit function
	// fmt.Println("\nTesting maxProfit:")
	// fmt.Println(maxProfit([]int{7,1,5,3,6,4})) // Should print: 5
	// fmt.Println(maxProfit([]int{7,6,4,3,1})) // Should print: 0
	// fmt.Println(maxProfit([]int{1,2,3,4,5})) // Should print: 4

	// Test isValid function
	// fmt.Println("\nTesting isValid:")
	// fmt.Println(isValid("()"))     // Should print: true
	// fmt.Println(isValid("()[]{}")) // Should print: true
	// fmt.Println(isValid("(]"))     // Should print: false
	// fmt.Println(isValid("([)]"))   // Should print: false
	// fmt.Println(isValid("{[]}"))   // Should print: true
	// Test minSubArrayLen function
// 	fmt.Println("\nTesting minSubArrayLen:")
// 	fmt.Println(minSubArrayLen(7, []int{2,3,1,2,4,3})) // Should print: 2
// 	fmt.Println(minSubArrayLen(4, []int{1,4,4})) // Should print: 1
// 	fmt.Println(minSubArrayLen(11, []int{1,1,1,1,1,1,1,1})) // Should print: 0
	// Test threeSum function
	fmt.Println("\nTesting threeSum:")
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // Should find triplets that sum to 0
	fmt.Println(threeSum([]int{0, 1, 1})) // Should print: []
	fmt.Println(threeSum([]int{0, 0, 0})) // Should print: [[0, 0, 0]]

}

