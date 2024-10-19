package main

import (
    "container/heap"
    "fmt"
    "math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } 
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func maxKelements(nums []int, k int) int64 {
    h := IntHeap(nums)
    heap.Init(&h)

    var score int64
    for i := 0; i < k; i++ {
        num := heap.Pop(&h).(int)

        score += int64(num)

        heap.Push(&h, int(math.Ceil(float64(num) / 3.0)))
    }
    return score
}

func main() {
    n := []int{1, 10, 3, 3, 3}
    fmt.Println(maxKelements(n, 3))  // Output should be the sum of the largest k elements, adjusted
}
