package main

import "fmt"


// brute force solution 
func countMaxOrSubsets(nums []int) int {
    max_or := 0
    for _, v := range nums{
        max_or |= v
    }
    var dfs func(int, int) int
    dfs = func(i int, curr_or int) int {
        if i == len(nums) {
            if curr_or == max_or{
                return 1
            }else {
                return 0
            }
        }

        return dfs(i+1 ,curr_or) + dfs(i +1,curr_or | nums[i])  
    }
    return dfs(0, 0)

}


func main() {
    fmt.Println("Hello, World!")
    nums := []int{3, 1}
    result := countMaxOrSubsets(nums)
    fmt.Println("Result:", result)
}
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
