package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF466D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7

	var n, h, v, cntL int64
	ans := int64(1)
	for Fscan(in, &n, &h); n > 0; n-- {
		Fscan(in, &v)
		v = h - v
		if v == cntL {
			ans = ans * (cntL + 1) % mod
		} else if v == cntL-1 {
			ans = ans * cntL % mod
		} else if v != cntL+1 {
			Fprint(out, 0)
			return
		}
		cntL = v
	}
	if cntL > 1 {
		ans = 0
	}
	Fprint(out, ans)
}

//func main() { CF466D(os.Stdin, os.Stdout) }
