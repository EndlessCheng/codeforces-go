package main

// 枚举所有三位数偶数

// github.com/EndlessCheng/codeforces-go
func findEvenNumbers(digits []int) (ans []int) {
	cnt := [10]int{}
	for _, d := range digits {
		cnt[d]++
	}
next:
	for i := 100; i < 1000; i += 2 { // 枚举所有三位数偶数 i
		c := [10]int{}
		for x := i; x > 0; x /= 10 { // 枚举 i 的每一位 d
			d := x % 10
			if c[d]++; c[d] > cnt[d] { // 如果 d 出现次数比 digits 中的 d 的次数还多，那么 i 肯定不能由 digits 中的数字组成
				continue next // 枚举下一个偶数
			}
		}
		ans = append(ans, i)
	}
	return
}
