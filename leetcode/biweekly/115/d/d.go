package main

// https://space.bilibili.com/206214
func countSubMultisets(nums []int, l, r int) (ans int) {
	const mod = 1_000_000_007
	total := 0
	cnt := map[int]int{}
	for _, x := range nums {
		total += x
		cnt[x]++
	}
	if l > total {
		return
	}

	r = min(r, total)
	f := make([]int, r+1)
	f[0] = cnt[0] + 1
	delete(cnt, 0)

	sum := 0
	for x, c := range cnt {
		sum = min(sum+x*c, r)
		for j := x; j <= sum; j++ {
			f[j] = (f[j] + f[j-x]) % mod // 同余前缀和
		}
		for j := sum; j >= x*(c+1); j-- {
			f[j] = (f[j] - f[j-x*(c+1)]) % mod
		}
	}

	for _, v := range f[l:] {
		ans += v
	}
	return (ans%mod + mod) % mod // 调整成非负数
}

func min(a, b int) int { if b < a { return b }; return a }
