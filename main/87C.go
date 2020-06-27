package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF87C(in io.Reader, out io.Writer) {
	n := 0
	Fscan(in, &n)

	type pair struct{ a0, len int }
	const mx = 1e5
	ps := [mx + 1][]pair{}
	for l := 2; l*(l+1)/2 <= mx; l++ {
		for a0 := 1; ; a0++ {
			s := (2*a0 + l - 1) * l / 2
			if s > mx {
				break
			}
			ps[s] = append(ps[s], pair{a0, l})
		}
	}
	sg := [mx + 1]int{}
	for i := range sg {
		mex := map[int]bool{}
		for _, p := range ps[i] {
			s := 0
			for j := p.a0; j-p.a0 < p.len; j++ {
				s ^= sg[j]
			}
			if i == n && s == 0 {
				Fprint(out, p.len)
				return
			}
			mex[s] = true
		}
		for ; mex[sg[i]]; sg[i]++ {
		}
	}
	Fprint(out, -1)
}

//func main() { CF87C(os.Stdin, os.Stdout) }
