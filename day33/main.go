package main

import "fmt"


func longestSquareStreak(nums []int) int {
    m := make(map[int]int)
    for _, v := range nums {
        if _, ok := m[v] ; !ok {
            m[v] = v*v
        }

    }
    max := 0
    for _, v := range nums {
        count := 0
        f := true
        n := v
        for f {
            if c , ok := m[n]; ok {
                n = c
                count++
            }else {
                max = maximum(max, count)
                
                f = false
            }
        }

    }
    if max == 1 {
        return -1
    }else {
        return max
    } 
    
}

func maximum(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{4,3,6,16,8,2}
    fmt.Println(longestSquareStreak(nums))
}
