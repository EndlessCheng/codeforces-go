package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func f17(v, w uint) (c int) {
	for v != w {
		if v < w {
			v, w = w, v
		}
		v = 1<<bits.Len(v-1) - v
		c++
	}
	return
}

func CF1617E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, p, q int
	Fscan(in, &n)
	a := make([]uint, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	mx := -1
	for i := 1; i < n; i++ {
		c := f17(a[i], a[0])
		if c > mx {
			mx, p = c, i
		}
	}

	mx = -1
	for i, v := range a {
		if i != p {
			c := f17(v, a[p])
			if c > mx {
				mx, q = c, i
			}
		}
	}
	Fprint(out, p+1, q+1, mx)
}

//func main() { CF1617E(os.Stdin, os.Stdout) }
