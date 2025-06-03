package main

import "fmt"

func canPartition(nums []int) bool {
	n := len(nums)
	totalSum := 0
	for k := range n {
		totalSum += nums[k]
	}
	if totalSum%2 != 0 {
		return false
	}
	target := totalSum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := range n {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]

}
func main() {
	arr := []int{1, 5, 11, 5}
	fmt.Println(canPartition(arr))
}
