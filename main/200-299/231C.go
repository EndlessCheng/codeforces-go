package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF231C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, l, mxC, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	s := int64(0)
	for r, v := range a {
		s += int64(v - a[0])
		for ; int64(v-a[0])*int64(r-l+1)-s > int64(k); l++ {
			s -= int64(a[l] - a[0])
		}
		if r-l+1 > mxC {
			mxC, ans = r-l+1, v
		}
	}
	Fprint(out, mxC, ans)
}

//func main() { CF231C(os.Stdin, os.Stdout) }
