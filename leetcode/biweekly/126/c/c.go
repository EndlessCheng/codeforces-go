package main

import (
	"container/heap"
	"math"
	"slices"
	"strings"
)

// https://space.bilibili.com/206214
func minimizeStringValue(s string) string {
	freq := [27]int{26: math.MaxInt / 26} // 哨兵
	for _, c := range s {
		if c != '?' {
			freq[c-'a']++
		}
	}
	f := freq
	slices.Sort(f[:])

	var limit, extra int
	q := strings.Count(s, "?")
	for i := 1; ; i++ {
		sum := i * (f[i] - f[i-1])
		if q <= sum {
			limit, extra = f[i-1]+q/i, q%i
			break
		}
		q -= sum
	}

	target := freq
	for i, c := range freq[:26] {
		if c > limit {
			continue
		}
		target[i] = limit
		if extra > 0 {
			extra--
			target[i]++
		}
	}

	ans := []byte(s)
	j := byte(0)
	for i, c := range ans {
		if c != '?' {
			continue
		}
		for freq[j] == target[j] {
			j++
		}
		freq[j]++
		ans[i] = 'a' + j
	}
	return string(ans)
}

//

func minimizeStringValue2(s string) string {
	h := make(hp, 26)
	for i := byte(0); i < 26; i++ {
		h[i].c = 'a' + i
	}
	for _, b := range s {
		if b != '?' {
			h[b-'a'].f++
		}
	}
	heap.Init(&h)

	t := make([]byte, strings.Count(s, "?"))
	for i := range t {
		t[i] = h[0].c
		h[0].f++
		heap.Fix(&h, 0)
	}
	slices.Sort(t)

	ans := []byte(s)
	j := 0
	for i, b := range ans {
		if b == '?' {
			ans[i] = t[j]
			j++
		}
	}
	return string(ans)
}

type pair struct {
	f int
	c byte
}
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.f < b.f || a.f == b.f && a.c < b.c }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
