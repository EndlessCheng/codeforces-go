package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1480B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, atk, hp, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &atk, &hp, &n)
		mxI := 0
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
			if a[i] > a[mxI] {
				mxI = i
			}
		}
		b := make([]int64, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		for i, v := range a {
			if i != mxI {
				hp -= ((b[i]-1)/atk + 1) * v
			}
		}
		hp -= (b[mxI] - 1) / atk * a[mxI]
		if hp > 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1480B(os.Stdin, os.Stdout) }
