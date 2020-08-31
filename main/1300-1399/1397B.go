package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1397B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}

	n, ans := 0, int64(1e18)
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
o:
	for c := int64(1); ; c++ {
		var s, p int64 = 0, 1
		for _, v := range a {
			if p > 1e10 {
				break o
			}
			s += abs(int64(v) - p)
			p *= c
		}
		if s < ans {
			ans = s
		}
	}
	Fprint(out, ans)
}

//func main() { CF1397B(os.Stdin, os.Stdout) }
