package main

// github.com/EndlessCheng/codeforces-go
func check(n int, cnt [3]int) bool {
	if cnt[1] == 0 {
		return false
	}
	cnt[1]--
	// 第一回合 Alice 移除 1，后面两人交替移除 1 和 2，中途可以插入 cnt[0] 个 0
	round := 1 + min(cnt[1], cnt[2])*2 + cnt[0]
	if cnt[1] > cnt[2] { // 可以再移除一个 1
		round++
	}
	return round < n && round%2 > 0
}

func stoneGameIX(stones []int) bool {
	cnt := [3]int{}
	for _, x := range stones {
		cnt[x%3]++
	}

	n := len(stones)
	// 小技巧：交换 cnt[1] 和 cnt[2] 再调用 check，相当于 Alice 第一回合移除了 2
	return check(n, cnt) || check(n, [3]int{cnt[0], cnt[2], cnt[1]})
}

func stoneGameIX2(stones []int) bool {
	cnt := [3]int{}
	for _, x := range stones {
		cnt[x%3]++
	}

	if cnt[0]%2 == 0 {
		return cnt[1] > 0 && cnt[2] > 0
	}
	return abs(cnt[1]-cnt[2]) > 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
