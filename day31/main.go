package main

import "fmt"


// daily 21/10/2024
func maxUniqueSplit(s string) int {
    set := make(map[string]int)

    l,r := 0,1
    for l < len(s) && r < len(s){
        fmt.Println(s[l:r])
        if _,ok := set[s[l:r]]; !ok {
            set[s[l:r]] = 1
            l = r
            r++
        }else {
            r++
        }
        fmt.Println(l,r)    
    }
    if _,ok := set[s[l:r]]; !ok {
            set[s[l:r]] = 1
    }
    fmt.Println(l,r)
    fmt.Println(set)

    return len(set)
}

func main() {
    fmt.Println(maxUniqueSplit("wwwzfvedwfvhsww"))
}
