package main

import "fmt"

func maximumSwap(num int) int {
    digits := []int{}
    for num > 0 {
        digits = append([]int{num % 10}, digits...)
        num /= 10
    }

    maxNum := digitsToNumber(digits)
    for i := 0; i < len(digits); i++ {
        for j := i + 1; j < len(digits); j++ {
            if digits[j] > digits[i] {
                swappedDigits := append([]int(nil), digits...)
                swappedDigits[j], swappedDigits[i] = swappedDigits[i], swappedDigits[j]
                maxNum = max(maxNum, digitsToNumber(swappedDigits))
            }
        }
    }

    return maxNum
}

func digitsToNumber(digits []int) int {
    number := 0
    for _, digit := range digits {
        number = number*10 + digit
    }
    return number
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(maximumSwap(1234123))
}
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
