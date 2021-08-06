package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF850C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	mask := map[int]int{}
	var n, x, xor int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x)
		for i := 2; i*i <= x; i++ {
			e := 0
			for ; x%i == 0; x /= i {
				e++
			}
			if e > 0 {
				mask[i] |= 1 << e
			}
		}
		if x > 1 {
			mask[x] |= 2
		}
	}

	SG := map[int]int{}
	var sg func(int) int
	sg = func(m int) (mex int) {
		m &^= 1
		if m == 0 {
			return
		}
		if v, has := SG[m]; has {
			return v
		}
		defer func() { SG[m] = mex }()
		has := map[int]bool{}
		for k := 1; 1<<k <= m; k++ {
			has[sg(m&(1<<k-1)|m>>k)] = true
		}
		for ; has[mex]; mex++ {
		}
		return
	}
	for _, m := range mask {
		xor ^= sg(m)
	}
	if xor > 0 {
		Fprint(out, "Mojtaba")
	} else {
		Fprint(out, "Arpa")
	}
}

//func main() { CF850C(os.Stdin, os.Stdout) }
