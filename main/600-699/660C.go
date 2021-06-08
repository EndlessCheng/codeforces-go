package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF660C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, l, c0, ans, posL int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for r := range a {
		Fscan(in, &a[r])
		c0 += a[r] ^ 1
		if c0 > k {
			for ; a[l] > 0; l++ {
			}
			l++
			c0--
		}
		if r-l+1 > ans {
			ans, posL = r-l+1, l
		}
	}
	Fprintln(out, ans)
	for i, v := range a {
		if posL <= i && i < posL+ans {
			v = 1
		}
		Fprint(out, v, " ")
	}
}

//func main() { CF660C(os.Stdin, os.Stdout) }
