package __99

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF87C(_r io.Reader, _w io.Writer) {
	type pair struct{ a1, len int }
	const mx = 1e5 + 5
	ps := [mx][]pair{}
	sgSum := [mx]int{}

	var n int
	Scan(&n)

	for len := 2; (len+1)*len/2 <= n; len++ {
		for a1 := 1; ; a1++ {
			s := (2*a1 + len - 1) * len / 2
			if s > n {
				break
			}
			ps[s] = append(ps[s], pair{a1, len})
		}
	}

	for i := 3; i <= n; i++ {
		mex := map[int]bool{}
		for _, p := range ps[i] {
			s := sgSum[p.a1] ^ sgSum[p.a1+p.len]
			if i == n && s == 0 {
				Print(p.len)
				return
			}
			mex[s] = true
		}
		sg := 0
		for ; mex[sg]; sg++ {
		}
		sgSum[i+1] = sgSum[i] ^ sg // SG 前缀和
	}
	Print(-1)
}
