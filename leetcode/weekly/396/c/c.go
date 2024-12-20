package main

// https://space.bilibili.com/206214
func minAnagramLength(s string) int {
	n := len(s)
	cntAll := [26]int{}
	for _, c := range s {
		cntAll[c-'a']++
	}
	g := 0
	for _, c := range cntAll {
		g = gcd(g, c)
	}
next:
	for times := g; times > 1; times-- {
		if g%times > 0 {
			continue
		}
		k := n / times
		cnt0 := [26]int{}
		for _, b := range s[:k] {
			cnt0[b-'a']++
		}
		for i := k * 2; i <= len(s); i += k {
			cnt := [26]int{}
			for _, b := range s[i-k : i] {
				cnt[b-'a']++
			}
			if cnt != cnt0 {
				continue next
			}
		}
		return k
	}
	return n
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }

func minAnagramLength2(s string) int {
	n := len(s)
next:
	for k := 1; k <= n/2; k++ {
		if n%k > 0 {
			continue
		}
		cnt0 := [26]int{}
		for _, b := range s[:k] {
			cnt0[b-'a']++
		}
		for i := k * 2; i <= len(s); i += k {
			cnt := [26]int{}
			for _, b := range s[i-k : i] {
				cnt[b-'a']++
			}
			if cnt != cnt0 {
				continue next
			}
		}
		return k
	}
	return n
}
