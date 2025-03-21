package main

// github.com/EndlessCheng/codeforces-go
const mx = 101

var divisors [mx][]int

func init() {
	// 预处理每个数的因子
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func countPairs(nums []int, k int) (ans int) {
	type pair struct{ v, d int }
	cnt := map[pair]int{}
	for j, x := range nums { // 枚举 j，计算左边有多少个符合要求的 i
		if j > 0 && x == nums[0] {
			ans++ // 单独统计 i=0 的情况
		}
		k2 := k / gcd(k, j) // i 必须是 k2 的倍数
		ans += cnt[pair{x, k2}] // 统计左边有多少个数，值为 x 且下标是 k2 的倍数
		for _, d := range divisors[j] {
			cnt[pair{x, d}]++ // j 是 d 的倍数
		}
	}
	return
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
