package main

import "fmt"


func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[rune]int)
    for i := 0; i < len(magazine); i++ {
        m[rune(magazine[i])] += 1

    }
    for i := 0; i < len(ransomNote); i++ {
        if v, ok := m[rune(ransomNote[i])]; ok && v > 0 {
            m[rune(ransomNote[i])] -= 1
        }else {
            return false
        }
    }
    return true
}


func main() {
    fmt.Println(canConstruct("aa", "ab"))
}
