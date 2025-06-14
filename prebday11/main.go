package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 2566. Maximum Difference by Remapping a Digit
// get the max by looping through the number looking fir the first non 9 digit and change all accunrce
// get the min by looping through the first
func minMaxDifference(num int) int {
    numStr := strconv.Itoa(num)
    maxN := math.MinInt
    minN := math.MaxInt

    for d := '0'; d <= '9'; d++ {
        replacedMax := strings.ReplaceAll(numStr, string(d), "9")
        valMax, _ := strconv.Atoi(replacedMax)
        if valMax > maxN {
            maxN = valMax
        }

        replacedMin := strings.ReplaceAll(numStr, string(d), "0")
        valMin, _ := strconv.Atoi(replacedMin)
        if valMin < minN {
            minN = valMin
        }
    }

    return maxN - minN
}

func main (){

	fmt.Println(minMaxDifference(11891))
}