package main

import (
    "fmt"
    "math/rand"
    "time"
)
// daily 30/9/20024 1381. Design a Stack With Increment Operation
type CustomStack struct {
    arr []int
	top int

}


func Constructor(maxSize int) CustomStack {
    return CustomStack{
		arr: make([]int, maxSize),
		top: 0,
	}
}


func (this *CustomStack) Push(x int)  {
    if this.top >= len(this.arr) {
		return
	}
	
	this.arr[this.top] = x
	this.top ++
}


func (this *CustomStack) Pop() int {
    if (this.top <= 0){
		return -1
	}
	val := this.arr[this.top -1]
	this.top --;
	return val
}


func (this *CustomStack) Increment(k int, val int)  {
    for i := 0; (i < k) && (i < len(this.arr)); i++ {
		this.arr[i] += val
	}
}


// func main() {
// 	stack := Constructor(3)
// 	stack.Push(1)
// 	stack.Push(2)
// 	fmt.Println(stack.Pop()) // Should print 2
// 	stack.Push(2)
// 	stack.Push(3)
// 	stack.Push(4) // This push should be ignored as the stack is full
// 	stack.Increment(5, 100)
// 	stack.Increment(2, 100)
// 	fmt.Println(stack.Pop()) // Should print 103
// 	fmt.Println(stack.Pop()) // Should print 202
// 	fmt.Println(stack.Pop()) // Should print 201
// 	fmt.Println(stack.Pop()) // Should print -1 as the stack is empty
// }\

// 380. Insert Delete GetRandom O(1)
type RandomizedSet struct {
	items map[int]int // to inset and delete in O(1)
	arr []int // to keep a count of numbers and length and get random
	top int
	size int
}


func Constructor2() RandomizedSet {
    return RandomizedSet{
		items: make(map[int]int),
		arr: make([]int, 200000),
		top: 0,
		size: 0,
	}
}


func (this *RandomizedSet) Insert(val int) bool {
    if _,e := this.items[val]; !e {
		this.items[val] = this.top;
		this.arr[this.top] = val ;
		this.top ++;
		this.size ++;
		return true
	}
	return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if _,e := this.items[val]; e {
		
		delete(this.items, val)
		this.size--;
		return true
	}
	return false
}


func (this *RandomizedSet) GetRandom() int {
    index := rand.Intn(this.top)
	if _,ok:= this.items[this.arr[index]]; ok {
		return this.arr[index]
	}else {
		return this.GetRandom()
	}
}


func main() {
	rand.Seed(time.Now().UnixNano())

	// Test RandomizedSet
	set := Constructor2()
	fmt.Println(set.Insert(1)) // Should print true
	fmt.Println(set.Remove(2)) // Should print false
	fmt.Println(set.Insert(2)) // Should print true
	fmt.Println(set.GetRandom()) // Should print either 1 or 2
	fmt.Println(set.Remove(1)) // Should print true
	fmt.Println(set.Insert(2)) // Should print false
	fmt.Println(set.GetRandom()) // Should print 2
}