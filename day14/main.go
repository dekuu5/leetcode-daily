package main

import (
	"fmt"
	
)

// //
// // daily 28/9/2024
// type Node struct {
// 	left *Node
// 	right *Node
// 	solts int
// 	start int
// 	end int
// }

// func CreateNode(start int, end int) *Node{
// 	return &Node{
// 			left: nil,
// 			right: nil,
// 			solts: 1,
// 			start: start,
// 			end: end,
// 		}
// }

// func (n *Node) insert(start int, end int) bool {
// 	curr := n
// 	for true {
// 		if (curr.start <= start || curr.end <= end  ) && curr.solts <=2 {
// 			curr = CreateNode(start, end)
// 			curr.solts++;
// 			return true
// 		}else if curr.end <= start {
// 			if curr.right == nil {
// 				curr.right = CreateNode(start, end)
// 				return true
// 			}
// 			curr = curr.right

// 		}else if curr.start >= end {
// 			if curr.left == nil {
// 				curr.left = CreateNode(start,end)
// 				return true
// 			}
// 			curr = curr.left
// 		}else {
// 			return false
// 		}
// 	}
// 	return false
// }

// type MyCalendar struct {
//     root *Node
// }

// func Constructor() MyCalendar {
//     return MyCalendar{
// 		root: nil,
// 	}
// }

// func (this *MyCalendar) Book(start int, end int) bool {
//     if this.root == nil {
// 		this.root = CreateNode(start, end)
// 		return true
// 	}
// 	return this.root.insert(start, end)
// }

// // ---------------------------------------------
// // top150 55. Jump Game
// // solving it like a two sum problem

// // func canJump(nums []int) bool {
// // 	Pi:=-1
// // 	curr:= 0
// // 	l:= len(nums)
// // 	for Pi != curr{
// // 		Pi = curr
// // 		curr= nums[curr] + curr
// // 		if curr >= l-1{
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// func main() {
// 	calendar := Constructor()
// 	fmt.Println(calendar.Book(10, 20)) // true
// 	fmt.Println(calendar.Book(50, 60)) // true
// 	fmt.Println(calendar.Book(10, 40)) // true
// 	fmt.Println(calendar.Book(5, 15))  // false
// 	fmt.Println(calendar.Book(5, 10))  // true
// 	fmt.Println(calendar.Book(25, 55)) // true
// }
// leaving this for another day

func reverseWords(s string) string {
	words := split(s)
	reverse(words)
	return join(words)
}
func split(words string) []string {
	w := []string{}
	startIndex := 0
	inWord := false

	for i := range words {
		if words[i] != ' ' {
			if !inWord {
				startIndex = i
				inWord = true
			}
		} else {
			if inWord {
				w = append(w, words[startIndex:i])
				inWord = false
			}
		}
	}
	if inWord {
		w = append(w, words[startIndex:])
	}
	return w
}

func reverse(words []string) {
	w := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		w[i] = words[len(words)-1-i]
	}
	copy(words, w)
}

func join(words []string) string{
	w := ""
	for i,word := range words{
		if i == len(words) -1 {
			w += word
		}else {
			w += word + " "
		}
	}
	return w
}


func main() {
	s := "  hello world  "
	
	fmt.Println(split(s))
	
	fmt.Println(reverseWords(s), "ww") // "blue is sky the"
}

