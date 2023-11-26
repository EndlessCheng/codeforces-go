package main

import "strings"

// https://space.bilibili.com/206214
func beautifulSubstrings(s string, k int) (ans int) {
	m := 1
	k *= 4
	for p := 2; p*p <= k; p++ {
		if k%p > 0 {
			continue
		}
		e := 1
		for k /= p; k%p == 0; k /= p {
			e++
		}
		m *= pow(p, (e+1)/2)
	}
	if k > 1 {
		m *= k
	}

	k = m

	sum := make([]int, len(s)+1)
	for i, v := range s {
		sum[i+1] = sum[i]
		if strings.Contains("aeiou", string(v)) {
			sum[i+1]++
		} else {
			sum[i+1]--
		}
	}

	pos := make([]map[int]int, len(sum)*2+10)
	for i := range pos {
		pos[i] = map[int]int{}
	}

	for i, s := range sum {
		s += len(sum)
		r := pos[s][i%k]
		ans += r
		pos[s][i%k]++
	}
	return
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x
		}
		x = x * x
	}
	return res
}
