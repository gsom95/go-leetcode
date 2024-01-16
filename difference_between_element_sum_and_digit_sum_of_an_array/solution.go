package differencebetweenelementsumanddigitsumofanarray

// src: https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/

func differenceOfSum(nums []int) int {
	elemSum := 0
	digitSum := 0
	for i, n := range nums {
		elemSum += nums[i]
		for n > 9 {
			digitSum += n % 10
			n /= 10
		}
		digitSum += n
	}

	return abs(elemSum - digitSum)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
