package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1483B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		next := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			next[i] = (i + 1) % n
		}

		ans := []interface{}{}
		q := []int{}
		for i, v := range a {
			if gcd(v, a[next[i]]) == 1 {
				q = append(q, i)
			}
		}
		vis := make([]bool, n)
		for len(q) > 0 {
			i := q[0]
			q = q[1:]
			if vis[i] {
				continue
			}
			if !vis[next[i]] {
				ans = append(ans, next[i]+1)
				vis[next[i]] = true
			}
			next[i] = next[next[i]]
			if gcd(a[i], a[next[i]]) == 1 {
				q = append(q, i)
			}
		}
		Fprint(out, len(ans), " ")
		Fprintln(out, ans...)
	}
}

//func main() { CF1483B(os.Stdin, os.Stdout) }
