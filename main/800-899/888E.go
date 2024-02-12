package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF888E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	b := a[:(n+1)/2]
	f := func(i int) int {
		s := int64(0)
		for j, v := range b {
			if i>>j&1 > 0 {
				s += v
			}
		}
		return int(s % int64(m))
	}
	x := []int{}
	for i := 0; i < 1<<len(b); i++ {
		s := f(i)
		ans = max(ans, s)
		x = append(x, s)
	}
	sort.Ints(x)
	b = a[(n+1)/2:]
	for i := 0; i < 1<<len(b); i++ {
		s := f(i)
		ans = max(ans, s)
		if j := sort.SearchInts(x, m-s); j > 0 {
			ans = max(ans, s+x[j-1])
		}
	}
	Fprint(out, ans)
}

//func main() { CF888E(os.Stdin, os.Stdout) }
