package main

import "fmt"

func minSwaps(s string) int {
    ans := 0 
    for i := 0; i < len(s); i++ {
        if s[i] == '['{
            ans++
        }else if ans > 0 {
            ans --
        }
    }
    return (ans+1)/2
}
// if i see [ it needs ] so ++ if you se [ 
// -- if you see ]

func main() {
    fmt.Println(minSwaps("][]["))
}
