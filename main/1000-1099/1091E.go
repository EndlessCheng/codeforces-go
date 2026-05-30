package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1091E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	slices.Sort(a[1:])

	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i]
	}

	l, r := 0, n
	for k, j := 1, 0; k <= n; k++ {
		for j < n && a[j+1] < k {
			j++
		}
		nk := n - k
		t := s[n] - s[n-k] - k*(k-1)
		if nk > 0 {
			c1 := min(nk, j)
			t -= s[c1]
			t -= k * (nk - c1)
		}
		l = max(l, t)
		r = min(r, a[n-k+1]+min(k, a[n-k+1])-t)
	}

	if l > r {
		Fprint(out, -1)
		return
	}
	for i := l; i <= r; i++ {
		if (s[n]+i)%2 == 0 {
			Fprint(out, i, " ")
		}
	}
}

//func main() { cf1091E(bufio.NewReader(os.Stdin), os.Stdout) }
