package main

import "fmt"

// func findKthBit(n int, k int) byte {
//     s := "0"
//     for i := 0; i < n; i++ {
//         s1 := reverse(invert(s))
//         s = s + "1" + s1
//     }
//     if k-1 < len(s) {
//         return s[k-1]
//     }
//     return '0'
// }

// func invert(s string) string {
//     var result string
//     for _, char := range s {
//         if char == '0' {
//             result += "1"
//         } else {
//             result += "0"
//         }
//     }
//     return result
// }

// func reverse(s string) string {
//     runes := []rune(s)
//     for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//         runes[i], runes[j] = runes[j], runes[i]
//     }
//     return string(runes)
// }


package main

import (
    "fmt"
)

func findKthBit(n int, k int) byte {
    return findKthBitRecursive(n, k)
}

func findKthBitRecursive(n int, k int) byte {
    // Base case: when n = 1, S(1) = "0"
    if n == 1 {
        return '0'
    }
    
    length := (1 << n) - 1 
    mid := length / 2 + 1 

    if k == mid {
        return '1' // Middle element is always "1"
    } else if k < mid {
        return findKthBitRecursive(n-1, k) // k is in the first half
    } else {
        mirrorPos := length - k + 1
        bit := findKthBitRecursive(n-1, mirrorPos)
        if bit == '0' {
            return '1'
        }
        return '0'
    }
}

func main() {
    n := 3
    k := 5
    fmt.Printf("The %d-th bit in S(%d) is: %c\n", k, n, findKthBit(n, k))
}


func main() {
    n := 3
    k := 5
    fmt.Printf("The %d-th bit in S(%d) is: %c\n", k, n, findKthBit(n, k))
}
