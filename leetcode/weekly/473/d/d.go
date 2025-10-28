package main

// https://space.bilibili.com/206214

// 974. 和可被 K 整除的子数组
// 由于本题 nums 没有负数，无需调整
func subarraysDivByK(nums []int, k int) (ans int64) {
	cnt := make(map[int]int, len(nums)) // 预分配空间
	s := 0
	for _, x := range nums {
		cnt[s]++
		s = (s + x) % k
		ans += int64(cnt[s])
	}
	return
}

func numGoodSubarrays(nums []int, k int) int64 {
	ans := subarraysDivByK(nums, k)
	start := 0
	for i, x := range nums {
		if i < len(nums)-1 && x == nums[i+1] {
			continue
		}
		// 遍历到了连续相同元素段的末尾
		size := i - start + 1 // 这一段的长度
		for sz := 1; sz <= size; sz++ {
			if x*sz%k == 0 {
				// 长为 sz 的子数组元素和能被 k 整除
				// 一共有 size-sz+1 个长为 sz 的子数组，其中有 size-sz 个重复的
				ans -= int64(size - sz)
			}
		}
		start = i + 1
	}
	return ans
}

func numGoodSubarrays2(nums []int, k int) (ans int64) {
	cnt := map[int]int{0: 1}
	sum := 0       // 前缀和
	lastStart := 0 // 上一个连续相同段的起始下标
	for i, x := range nums {
		if i > 0 && x != nums[i-1] {
			// 上一个连续相同段结束，可以把上一段对应的前缀和添加到 cnt
			s := sum
			for range i - lastStart {
				cnt[s%k]++
				s -= nums[i-1]
			}
			lastStart = i
		}
		sum += x
		ans += int64(cnt[sum%k])
	}
	return
}

func numGoodSubarrays1(nums []int, k int) (ans int64) {
	n := len(nums)
	cnt := map[int]int{0: 1}
	s := 0
	for i := 0; i < n; {
		start := i
		x := nums[start]
		s0 := s
		for ; i < n && nums[i] == x; i++ {
			s += x
			ans += int64(cnt[s%k])
		}
		for range i - start {
			s0 += x
			cnt[s0%k]++
		}
	}
	return
}
