package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF354C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx int = 1e6

	var n, lim, v int
	Fscan(in, &n, &lim)
	has := [mx + 1]bool{}
	mi := mx
	for ; n > 0; n-- {
		Fscan(in, &v)
		has[v] = true
		if v < mi {
			mi = v
		}
	}
	l := [mx * 2]int{}
	for i, cur := 1, -1; i < mx*2; i++ {
		if i <= mx && has[i] {
			cur = i
		}
		l[i] = cur
	}
o:
	for i := mi; ; i-- {
		for j := i*2 - 1; j < mx*2; j += i {
			if l[j]%i > lim {
				continue o
			}
		}
		Fprint(out, i)
		return
	}
}

//func main() { CF354C(os.Stdin, os.Stdout) }
