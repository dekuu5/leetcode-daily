package main

import "fmt"

func minimumSteps(s string) int64 {

    steps := 0
    l:= 0
    runes := []rune(s)
    for r,v := range runes {
        if v == '0' {
            steps += r-l
            l++

        }
        
        
    }
    s = string(runes)
    return int64(steps)
    
}


func main() {
    fmt.Println(minimumSteps("11000111"))
}
