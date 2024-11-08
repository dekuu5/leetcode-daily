package main

import (
	"fmt"
)

func minChanges(s string) int {
    res := 0
    for i := 0; i < len(s); i+=2 {
        
        if s[i] != s[i+1] {
            res+=1
        }
    }
    return res
}



func main() {
    fmt.Println(minChanges("01010001"))
}
