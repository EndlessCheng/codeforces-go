package main

// github.com/EndlessCheng/codeforces-go
func smallestSubsequence(s string) string {
	left := [26]int{}
	for _, b := range s {
		left[b-'a']++
	}
	vis := [26]bool{}
	stk := []byte{}
	for i := range s {
		b := s[i]
		if !vis[b-'a'] {
			for len(stk) > 0 && b < stk[len(stk)-1] {
				last := stk[len(stk)-1] - 'a'
				if left[last] == 0 {
					break
				}
				vis[last] = false
				stk = stk[:len(stk)-1]
			}
			vis[b-'a'] = true
			stk = append(stk, b)
		}
		left[b-'a']--
	}
	return string(stk)
}
