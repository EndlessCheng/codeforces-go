package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1042C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, mxI int
	Fscan(in, &n)
	var neg, zero, pos []int
	for i, mxV := 1, -int(2e9); i <= n; i++ {
		if Fscan(in, &v); v < 0 {
			if v > mxV {
				mxI, mxV = i, v
			}
			neg = append(neg, i)
		} else if v == 0 {
			zero = append(zero, i)
		} else {
			pos = append(pos, i)
		}
	}

	if m := len(neg); m&1 > 0 {
		zero = append(zero, mxI)
		for i, p := range neg {
			if p == mxI {
				pos = append(append(pos, neg[:i]...), neg[i+1:]...)
			}
		}
	} else {
		pos = append(pos, neg...)
	}

	if len(zero) > 1 {
		for _, p := range zero[1:] {
			Fprintln(out, 1, p, zero[0])
		}
	}
	if len(pos) > 0 {
		if len(zero) > 0 {
			Fprintln(out, 2, zero[0])
		}
		for _, p := range pos[1:] {
			Fprintln(out, 1, p, pos[0])
		}
	}
}

//func main() { CF1042C(os.Stdin, os.Stdout) }
