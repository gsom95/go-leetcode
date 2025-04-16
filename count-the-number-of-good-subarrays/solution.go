package countthenumberofgoodsubarrays

func countGood(nums []int, k int) int64 {
	nLen := len(nums)
	same, right := 0, -1
	cnt := make(map[int]int)
	ans := int64(0)
	for left := range nLen {
		for same < k && right+1 < nLen {
			right++
			same += cnt[nums[right]]
			cnt[nums[right]]++
		}
		if same >= k {
			ans += int64(nLen - right)
		}
		cnt[nums[left]]--
		same -= cnt[nums[left]]
	}
	return ans
}
