package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF925C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var v, sum uint64
	Fscan(in, &n)
	ans := make([]interface{}, n)
	a := make([][]uint64, 60)
	for ; n > 0; n-- {
		Fscan(in, &v)
		l := bits.Len64(v) - 1
		a[l] = append(a[l], v)
	}
o:
	for i := range ans {
		for j, b := range a {
			if len(b) > 0 && sum>>j&1 == 0 {
				sum ^= b[0]
				ans[i] = b[0]
				a[j] = a[j][1:]
				continue o
			}
		}
		Fprint(out, "No")
		return
	}
	Fprintln(out, "Yes")
	Fprintln(out, ans...)
}

//func main() { CF925C(os.Stdin, os.Stdout) }
