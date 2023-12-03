package main

// https://space.bilibili.com/206214
func f(s string, k int) (res int) {
	for m := 1; m <= 26 && k*m <= len(s); m++ {
		cnt := [26]int{}
		check := func() {
			for i := range cnt {
				if cnt[i] > 0 && cnt[i] != k {
					return
				}
			}
			res++
		}
		for right, in := range s {
			cnt[in-'a']++
			if left := right + 1 - k*m; left >= 0 {
				check()
				cnt[s[left]-'a']--
			}
		}
	}
	return
}

func countCompleteSubstrings(word string, k int) (ans int) {
	for i, n := 0, len(word); i < n; {
		st := i
		for i++; i < n && abs(int(word[i])-int(word[i-1])) <= 2; i++ {
		}
		ans += f(word[st:i], k)
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }