package main

import "fmt"

// brute force
// func maxWidthRamp(nums []int) int {
//     mWidth := 0
//     for i := 0; i < len(nums); i++ {
//         for j := i+1; j < len(nums); j++ {
//             if nums[i] <= nums[j]  && j-i > m{
//                 mWidth = j - i  
//                 println(i,j,mWidth)
//             }
//         }
//     }
//     return mWidth
// }

type Stack struct {
    items []int
    top   int
}

func Constructor() Stack {
    return Stack{
        items: make([]int, 0),
        top:   -1,
    }
}

func (s *Stack) Push(c int) {
    s.items = append(s.items, c)
    s.top++
}

func (s *Stack) Pop() int {
    if s.top == -1 {
        return 0
    }
    c := s.items[s.top]
    s.items = s.items[:s.top]
    s.top--
    return c
}

func (s *Stack) Peek() int {
    if s.top == -1 {
        return 0
    }
    return s.items[s.top]
}

func (s *Stack) Size() int {
    return s.top + 1
}

func (s *Stack) IsEmpty() bool {
    return s.top == -1
}

func maxWidthRamp(nums []int) int {
    stk := Constructor()

    for i := 0; i < len(nums); i++ {
        if stk.IsEmpty() || nums[stk.Peek()] > nums[i] {
            stk.Push(i)
        }
    }

    mWidth := 0

    for i := len(nums) - 1; i >= 0; i-- {
        for !stk.IsEmpty() && nums[stk.Peek()] <= nums[i] {
            mWidth = max(mWidth, i - stk.Peek())
            stk.Pop()
        }
    }
    
    return mWidth
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

 


func main() {
    fmt.Println(maxWidthRamp([]int{9,8,1,0,1,9,4,0,4,1}))
}
