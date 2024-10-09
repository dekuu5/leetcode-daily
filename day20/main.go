package main

import "fmt"



// func findMin(s string) string {
//     for i := 0; i < len(s)-1; i++ { 
//         if s[i] == 'A' && s[i+1] == 'B' {
//             newS := ""
//             if i+2 > len(s)-1 {
//                 newS = s[:i]
//             } else {
//                 newS = s[:i] + s[i+2:]
//             }
//             return findMin(newS)
//         } else if s[i] == 'C' && s[i+1] == 'D' {
//             newS := ""
//             if i+2 > len(s)-1 {
//                 newS = s[:i]
//             } else {
//                 newS = s[:i] + s[i+2:]
//             }
//             return findMin(newS)
//         }
//     }
//     return s
// }


type Stack struct {
    items []byte
    top   int
}

func Constructor() Stack {
    return Stack{
        items: make([]byte, 0),
        top:   -1,
    }
}

func (s *Stack) Push(c byte) {
    s.items = append(s.items, c)
    s.top++
}

func (s *Stack) Pop() (byte, bool) {
    if s.top == -1 {
        return ' ', false
    }
    c := s.items[s.top]
    s.items = s.items[:s.top] 
    s.top--                
    return c, true
}

func (s *Stack) Peek() (byte, bool) {
    if s.top == -1 { 
        return ' ', false
    }
    return s.items[s.top], true
}

func (s *Stack) Peek2() (byte, bool) {
    if s.top <= 0 {
        return ' ', false
    }
    return s.items[s.top-1], true
}

func (s *Stack) Size() int {
    return s.top + 1
}

func minLength(s string) int {
    stk := Constructor()

    for i := 0; i < len(s); i++ {
        stk.Push(s[i])
        if stk.Size() >= 2 {
            top, _ := stk.Peek()
            secondTop, _ := stk.Peek2()
            if (top == 'B' && secondTop == 'A') || (top == 'D' && secondTop == 'C') {
                stk.Pop()
                stk.Pop()
            }
        }
    }

    return stk.Size()
}

func main() {
    fmt.Println(minLength("ABFCACDB")) // Expected output: 2
}

