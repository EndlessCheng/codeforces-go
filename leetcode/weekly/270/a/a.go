package main

// github.com/EndlessCheng/codeforces-go
func findEvenNumbers(digits []int) (ans []int) {
	cnt := make([]int, 10)
	for _, d := range digits {
		cnt[d]++
	}

	// i=0 百位，i=1 十位，i=2 个位，x 表示当前正在构造的数字
	var dfs func(i, x int)
	dfs = func(i, x int) {
		if i == 3 {
			ans = append(ans, x)
			return
		}
		for d, c := range cnt {
			if c > 0 && (i == 0 && d > 0 || i == 1 || i == 2 && d%2 == 0) {
				cnt[d]--
				dfs(i+1, x*10+d)
				cnt[d]++ // 恢复现场
			}
		}
	}
	dfs(0, 0)
	return
}

func findEvenNumbers1(digits []int) (ans []int) {
	cnt := [10]int{}
	for _, d := range digits {
		cnt[d]++
	}

next:
	for i := 100; i < 1000; i += 2 { // 枚举所有三位数偶数 i
		c := [10]int{}
		for x := i; x > 0; x /= 10 { // 枚举 i 的每一位 d
			d := x % 10
			c[d]++
			if c[d] > cnt[d] { // 如果 i 中 d 的个数比 digits 中的还多，那么 i 无法由 digits 中的数字组成
				continue next // 枚举下一个偶数
			}
		}
		ans = append(ans, i)
	}
	return
}
