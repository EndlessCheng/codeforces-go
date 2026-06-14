package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf549F(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		s[i] = (s[i-1] + a[i]) % k
	}

	pos := make([][]int, k)
	for i, v := range s {
		pos[v] = append(pos[v], i)
	}
	ask := func(v, L, R int) int { return sort.SearchInts(pos[v], R+1) - sort.SearchInts(pos[v], L) }

	l := make([]int, n+1)
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		l[i] = i - 1
		for l[i] > 0 && a[l[i]] < a[i] {
			l[i] = l[l[i]]
		}
	}
	for i := n; i > 0; i-- {
		r[i] = i + 1
		for r[i] <= n && a[r[i]] <= a[i] {
			r[i] = r[r[i]]
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if i-l[i] < r[i]-i {
			for j := l[i] + 1; j <= i; j++ {
				v := (a[i] + s[j-1]) % k
				ans += ask(v, i, r[i]-1)
			}
		} else {
			for j := i; j <= r[i]-1; j++ {
				v := (s[j] - a[i]%k + k) % k
				ans += ask(v, l[i], i-1)
			}
		}
	}
	Fprint(out, ans-n)
}

//func main() { cf549F(bufio.NewReader(os.Stdin), os.Stdout) }
