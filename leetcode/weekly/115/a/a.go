package main

// github.com/EndlessCheng/codeforces-go
func prisonAfterNDays(cells []int, n int) (ans []int) {
	mp := map[int]int{}
	s := 0
	for i, cell := range cells {
		s |= cell << i
	}
o:
	for i := 0; i < n; i++ {
		if st, has := mp[s]; has {
			n = (n-st)%(i-st) + st
			for t, d := range mp {
				if d == n {
					s = t
					break o
				}
			}
		}
		mp[s] = i
		t := 0
		for j := 1; j < 7; j++ {
			if s>>(j-1)&1 == s>>(j+1)&1 {
				t |= 1 << j
			}
		}
		s = t
	}
	for i := 0; i < 8; i++ {
		ans = append(ans, s>>i&1)
	}
	return
}
