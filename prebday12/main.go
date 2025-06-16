package main


func maximumDifference(nums []int) int {
    maxDiff := -1

	minN := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] <= minN {
			minN = nums[i]
			continue
		} else {
			maxDiff = max(maxDiff , nums[i] - minN)
		}
		
	}
	return maxDiff
}


func main(){

}