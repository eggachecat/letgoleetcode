package search

import "log"

type SolutionsFor53 interface {
	MaxSubArray(nums []int) int
}

type SolutionsFor53At0 struct{}
type SolutionsFor53At1 struct{}
type SolutionsFor53At2 struct{}

// 一种子问题 f(k):
// 		nums[1],nums[2],...,nums[k]
//	似乎会碰到一个困难: 如果只考虑nums[k]在不在解中,显然会有问题(例如真正的解是从nums[k]开始)
// 所以可能需要 f(i,j)
//  	nums[i],nums[i+1],...,nums[j]
// 就把他们和全部求出来,然后找到最大即可
func (s *SolutionsFor53At0) MaxSubArray(nums []int) int {
	// SolutionsFor53At0 最容易想到 但是内存爆了
	// -4 -2 -1 -3
	//  0  2  3  2
	//  0  0  1 -1
	//  0  0  0 -2
	// 我们可以看到似乎可以复写上去
	nNums := len(nums)

	// sums[i][j] 标识 nums[i]+...+nums[j]
	sums := make([][]int, nNums)
	maxSum := nums[0]
	for i := 0; i < nNums; i++ {
		sums[i] = make([]int, nNums)
	}
	sums[0][0] = nums[0]
	for i := 0; i < nNums; i++ {
		// 看到第i个数
		sums[i][i] = nums[i]
		if sums[i][i] > maxSum {
			maxSum = sums[i][i]
		}
		for j := 0; j < i; j++ {
			// nums[j][i]都做个更新： (nums[j]+..+nums[i-1]) + nums[i]
			sums[j][i] = sums[j][i-1] + nums[i]
			if sums[j][i] > maxSum {
				maxSum = sums[j][i]
			}
		}
	}
	log.Println(sums)
	return maxSum
}

func (s *SolutionsFor53At1) MaxSubArray(nums []int) int {
	// SolutionsFor53At0 复写上去
	nNums := len(nums)

	sums := make([]int, nNums)
	maxSum := nums[0]

	for i := 0; i < nNums; i++ {
		// 看到第i个数
		sums[i] = nums[i]
		if sums[i] > maxSum {
			maxSum = sums[i]
		}
		for j := 0; j < i; j++ {
			// nums[j][i]都做个更新： (nums[j]+..+nums[i-1]) + nums[i]
			sums[j] = sums[j] + nums[i]
			if sums[j] > maxSum {
				maxSum = sums[j]
			}
		}
	}
	log.Println(sums)
	return maxSum
}

// 显然子问题我们可以限制成:
//		f(k)
//         如果subarray终止再[k]!
//			两选择在i: i-1的加上nums[i] 或者只选 i
func (s *SolutionsFor53At2) MaxSubArray(nums []int) int {
	nNums := len(nums)
	subArraySums := make([]int, nNums)
	maxSum := nums[0]
	subArraySums[0] = nums[0]

	for i := 1; i < nNums; i++ {
		choice0 := nums[i]
		// 显然不用全部记下来就记录上一个就好了
		choice1 := subArraySums[i-1] + nums[i]

		if choice0 < choice1 {
			subArraySums[i] = choice1
		} else {
			subArraySums[i] = choice0
		}

		if subArraySums[i] > maxSum {
			maxSum = subArraySums[i]
		}

	}
	return maxSum
}
