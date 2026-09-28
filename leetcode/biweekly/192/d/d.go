package main

import (
	"sort"
)

// https://space.bilibili.com/206214
func longestSubarray1(nums []int, k int) (ans int) {
	numPos := make([][]int, k) // 2*nums[i] % k 出现的所有位置
	firstPos := make([]int, k) // 前缀和 % k 首次出现的位置
	for i := range firstPos {
		firstPos[i] = -1
	}
	lastPos := make([]int, k) // 前缀和 % k 最后一次出现的位置
	firstPos[0] = 0
	lastPos[0] = 0
	sum := 0 // 前缀和

	for i, x := range nums {
		x = x%k + k // 保证 x 非负
		x2 := x * 2 % k
		numPos[x2] = append(numPos[x2], i)

		sum = (sum + x) % k
		r := i + 1
		l := firstPos[sum]
		if l < 0 {
			firstPos[sum] = r
		} else {
			// 不取反时的最大长度
			ans = max(ans, r-l)
		}
		lastPos[sum] = r
	}

	// 枚举 s[l]%k 和 s[r]%k，判断是否存在满足要求的 i
	for sl, l := range firstPos {
		if l < 0 {
			continue
		}
		for sr, r := range lastPos {
			if r-l <= ans { // 最优性优化：ans 无法增大
				continue
			}
			// 2*nums[i]%k = (sr-sl)%k
			pos := numPos[(sr-sl+k)%k] // +k 保证结果非负
			idx := sort.SearchInts(pos, l)
			if idx < len(pos) && pos[idx] < r {
				ans = r - l
			}
		}
	}
	return
}

func longestSubarray(nums []int, k int) (ans int) {
	firstSum := []int{0}
	firstPos := make([]int, k) // 前缀和 % k 首次出现的位置
	for i := range firstPos {
		firstPos[i] = -1
	}
	lastPos := make([]int, k) // 前缀和 % k 最后一次出现的位置
	firstPos[0] = 0
	lastPos[0] = 0
	sum := 0 // 前缀和

	for i, x := range nums {
		x = x%k + k // 保证 x 非负
		nums[i] = x

		sum = (sum + x) % k
		r := i + 1
		l := firstPos[sum]
		if l < 0 {
			firstPos[sum] = r
			firstSum = append(firstSum, sum)
		} else {
			// 不取反时的最大长度
			ans = max(ans, r-l)
		}
		lastPos[sum] = r
	}

	lastX2 := make([]int, k)
	for i := range lastX2 {
		lastX2[i] = -1
	}
	sr := 0

	// 枚举 s[r]%k 和 s[l]%k，判断是否存在满足要求的 i
	for i, x := range nums {
		lastX2[x*2%k] = i
		sr = (sr + x) % k
		r := i + 1
		if lastPos[sr] != r { // 只考虑 s[r]%k 最后一次出现的位置
			continue
		}
		for _, sl := range firstSum {
			l := firstPos[sl]
			if r-l <= ans { // 最优性优化：ans 无法增大
				break
			}
			j := lastX2[(sr-sl+k)%k] // +k 保证结果非负
			if j >= l {
				ans = r - l
			}
		}
	}

	return
}

func longestSubarray3x(nums []int, k int) (ans int) {
	type pair struct{ l, sum int }
	firsts := []pair{{}}
	firstPos := make([]int, k) // 前缀和 % k 首次出现的位置
	for i := range firstPos {
		firstPos[i] = -1
	}
	lastPos := make([]int, k) // 前缀和 % k 最后一次出现的位置
	firstPos[0] = 0
	lastPos[0] = 0
	sum := 0 // 前缀和

	for i, x := range nums {
		x = x%k + k // 保证 x 非负
		nums[i] = x

		sum = (sum + x) % k
		r := i + 1
		l := firstPos[sum]
		if l < 0 {
			firstPos[sum] = r
			firsts = append(firsts, pair{r, sum})
		} else {
			// 不取反时的最大长度
			ans = max(ans, r-l)
		}
		lastPos[sum] = r
	}

	// ptr[2*nums[i]%k] 表示 2*nums[i]%k 目前处理到的 firsts 的下标
	ptr := make([]int, k)

	// 枚举 2*nums[i]%k 和 s[l]%k，判断是否存在满足要求的 r
	for i, x := range nums {
		x2 := x * 2 % k
		p := &ptr[x2]
		for *p < len(firsts) && firsts[*p].l <= i {
			r := lastPos[(firsts[*p].sum+x2)%k]
			if i < r {
				ans = max(ans, r-firsts[*p].l)
			}
			// 后面遇到相同的 x2，可以从 ptr[x2] 开始继续枚举，从而避免重复枚举
			*p++
		}
	}

	return
}
