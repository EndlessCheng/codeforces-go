package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1278C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, 2*n)
		sum := 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == 2 {
				a[i] = -1
			}
			sum += a[i]
		}
		if sum == 0 {
			Fprintln(out, 0)
			continue
		}
		ans := 2 * n
		l := map[int]int{0: 0}
		s := 0
		for i, v := range a[n:] {
			i++
			s += v
			if _, ok := l[s]; !ok {
				l[s] = i
			}
			if s == sum && i < ans { // 注：写的时候没考虑全取右侧的情况，WA 了两发
				ans = i
			}
		}
		for i := n - 1; i >= 0; i-- {
			sum -= a[i]
			if l, ok := l[sum]; ok && n-i+l < ans {
				ans = n - i + l
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1278C(os.Stdin, os.Stdout) }
