package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1111B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	var n, k, m int64
	Fscan(in, &n, &k, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	s := int64(0)
	for _, v := range a {
		s += int64(v)
	}
	ans := float64(s) / float64(n)
	for i := int64(0); i <= min(n-1, m); i++ {
		ans = math.Max(ans, float64(s+min(m-i, k*(n-i)))/float64(n-i))
		s -= int64(a[i])
	}
	Fprintf(out, "%.20f", ans)
}

//func main() { CF1111B(os.Stdin, os.Stdout) }
