package main

import "fmt"


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
func (s *Stack) IsEmpty() bool {
    return s.top == -1
}

func minAddToMakeValid(s string) int {
    stk := Constructor()
    steps :=0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stk.Push('(')
        }else if !stk.IsEmpty(){
            stk.Pop()
        }else {
            steps++
        }
    }
    for !stk.IsEmpty() {
        steps++
        stk.Pop()
    }
    return steps
}


func main() {
    fmt.Println(minAddToMakeValid("((("))
}
