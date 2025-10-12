package main

// https://space.bilibili.com/206214
func longestBalanced1(s string) (ans int) {
	for i := range s {
		cnt := make([]int, 26)
	next:
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a']++
			base := cnt[s[j]-'a']
			for _, c := range cnt {
				if c > 0 && c != base {
					continue next
				}
			}
			ans = max(ans, j-i+1)
		}
	}
	return
}

func longestBalanced(s string) (ans int) {
	for i := range s {
		cnt := [26]int{}
		mx, kinds := 0, 0
		for j := i; j < len(s); j++ {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				kinds++
			}
			cnt[b]++
			mx = max(mx, cnt[b])
			if mx*kinds == j-i+1 {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}
