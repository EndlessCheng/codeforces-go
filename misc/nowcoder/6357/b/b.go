package main

// github.com/EndlessCheng/codeforces-go
func solve(S, T string) int {
	delta := [3]byte{2, 3, 5}
	var s, t [4]byte
	for i := range S {
		s[i] = S[i]
	}
	for i := range T {
		t[i] = T[i]
	}
	vis := map[[4]byte]bool{s: true}
	q := [][4]byte{s}
	for d := 0; ; d++ {
		qq := q
		q = nil
		for _, s := range qq {
			if s == t {
				return d
			}
			for i := 0; i < 4; i++ {
				c := 0
				ss := s
				for j := 0; j < 4; j++ {
					if j == i {
						continue
					}
					ss[j] += delta[c]
					if ss[j] > 'z' {
						ss[j] -= 26
					}
					c++
				}
				if !vis[ss] {
					vis[ss] = true
					q = append(q, ss)
				}
			}
		}
	}
}
