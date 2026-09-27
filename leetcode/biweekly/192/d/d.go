package main

import "sort"

// https://space.bilibili.com/206214
func longestSubarray0(nums []int, k int) (ans int) {
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
		y := x * 2 % k
		numPos[y] = append(numPos[y], i)

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

func longestSubarray1(nums []int, k int) (ans int) {
	// 记录前缀和 % k 最后一次出现的下标
	lastPos := make([]int, k)
	for i := range lastPos {
		lastPos[i] = -1
	}

	lastPos[0] = 0
	sum := 0
	for i, x := range nums {
		x = x%k + k // 保证 nums[i] 非负
		nums[i] = x
		sum = (sum + x) % k
		lastPos[sum] = i + 1
	}

	// 记录前缀和 % k 首次出现的下标
	firstPos := make([]int, k) // todo vis 优化
	for i := range firstPos {
		firstPos[i] = -1
	}

	visNum := make([]bool, k)
	firstPos[0] = 0
	sum = 0
	for i, x := range nums {
		sum = (sum + x) % k
		// 发现新的前缀和 % k
		if firstPos[sum] < 0 {
			firstPos[sum] = i + 1
			clear(visNum)
		} else {
			ans = max(ans, i+1-firstPos[sum])
		}

		y := x * 2 % k
		if visNum[y] {
			// 没有发现新的前缀和 % k，不考虑重复的 2x % k 
			continue
		}
		visNum[y] = true

		for sumL, l := range firstPos {
			if l < 0 || l == i+1 {
				continue
			}
			r := lastPos[(sumL+y)%k]
			if r > i {
				ans = max(ans, r-l)
			}
		}
	}

	return
}

func longestSubarray(nums []int, k int) (ans int) {
	// 记录前缀和 % k 最后一次出现的下标
	lastPos := make([]int, k)
	for i := range lastPos {
		lastPos[i] = -1
	}

	lastPos[0] = 0
	sum := 0
	for i, x := range nums {
		x = x%k + k
		nums[i] = x // 保证 nums[i] 非负
		sum = (sum + x) % k
		lastPos[sum] = i + 1
	}

	// 记录前缀和 % k 首次出现的下标
	firstPos := make([]int, k)
	for i := range firstPos {
		firstPos[i] = -1
	}
	type pair struct{ sum, l int }
	first := []pair{{}}

	visTime := make([]int, k)
	t := 0
	sum = 0
	for i, x := range nums {
		// 发现新的前缀和 % k
		if firstPos[sum] < 0 {
			firstPos[sum] = i
			first = append(first, pair{sum, i})
			t++
		}

		sum = (sum + x) % k
		l := firstPos[sum]
		if l >= 0 {
			ans = max(ans, i+1-l) // 不取反的情况
		}

		y := x * 2 % k
		if visTime[y] == t {
			// 没有发现新的前缀和 % k，不考虑重复的 2x % k 
			continue
		}
		visTime[y] = t

		for _, p := range first {
			r := lastPos[(p.sum+y)%k]
			if r > i {
				ans = max(ans, r-p.l)
			}
		}
	}

	return
}
