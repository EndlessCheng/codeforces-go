package main

// https://space.bilibili.com/206214
func pSqrt(n int) int {
	res := 1
	for i := 2; i*i <= n; i++ {
		i2 := i * i
		for n%i2 == 0 {
			res *= i
			n /= i2
		}
		if n%i == 0 {
			res *= i
			n /= i
		}
	}
	if n > 1 {
		res *= n
	}
	return res
}

func beautifulSubstrings(s string, k int) (ans int64) {
	k = pSqrt(k * 4)

	type pair struct{ i, sum int }
	cnt := map[pair]int{{k - 1, 0}: 1} // k-1 和 -1 同余
	sum := 0
	const aeiouMask = 1065233
	for i, c := range s {
		bit := aeiouMask >> (c - 'a') & 1
		sum += bit*2 - 1
		p := pair{i % k, sum}
		ans += int64(cnt[p])
		cnt[p]++
	}
	return
}
