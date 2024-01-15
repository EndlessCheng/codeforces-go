package main

// https://space.bilibili.com/206214
func beautifulIndices(s, a, b string, k int) (ans []int) {
	posA := kmp(s, a)
	posB := kmp(s, b)

	j, m := 0, len(posB)
	for _, i := range posA {
		for j < m && posB[j] < i-k {
			j++
		}
		if j < m && posB[j] <= i+k {
			ans = append(ans, i)
		}
	}
	return
}

func kmp(text, pattern string) (pos []int) {
	m := len(pattern)
	pi := make([]int, m)
	cnt := 0
	for i := 1; i < m; i++ {
		v := pattern[i]
		for cnt > 0 && pattern[cnt] != v {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == v {
			cnt++
		}
		pi[i] = cnt
	}

	cnt = 0
	for i, v := range text {
		for cnt > 0 && pattern[cnt] != byte(v) {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == byte(v) {
			cnt++
		}
		if cnt == m {
			pos = append(pos, i-m+1)
			cnt = pi[cnt-1]
		}
	}
	return
}
