package main

import "fmt"


func productExceptSelf(nums []int) []int {
    n:= len(nums)
    // arr := make([]int, n)
    prefix := make([]int, n)
    postfix := make([]int, n)

    prefix[0] = 1
    pre := nums[0]
    for i := 1; i < n; i++ {
        prefix[i] = pre
        pre *= nums[i]
    }
    postfix[n-1] = 1
    post := nums[n-1]
    for i := n-2; i >= 0 ; i-- {
        postfix[i] = post
        post*= nums[i]    
    }
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = postfix[i] * prefix[i]
    }

    return res

}
func main() {
    fmt.Println(productExceptSelf([]int{1,2,3,4}))
}
