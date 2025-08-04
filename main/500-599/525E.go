package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf525E(in io.Reader, out io.Writer) {
	var n, k, tar, ans int
	Fscan(in, &n, &k, &tar)
	F := []int{1}
	for i := 1; F[len(F)-1]*i <= tar; i++ {
		F = append(F, F[len(F)-1]*i)
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	calc := func(a []int) map[int][]int {
		cnt := map[int][]int{}
		var f func(int, int, int)
		f = func(i, j, s int) {
			if s > tar {
				return
			}
			if i < 0 {
				if _, ok := cnt[s]; !ok {
					cnt[s] = make([]int, k+1)
				}
				cnt[s][j]++
				return
			}
			f(i-1, j, s)
			v := a[i]
			f(i-1, j, s+v)
			if j < k && v < len(F) {
				f(i-1, j+1, s+F[v])
			}
		}
		f(len(a)-1, 0, 0)
		return cnt
	}

	cnt := calc(a[:n/2])
	for x, b := range calc(a[n/2:]) {
		if c, ok := cnt[tar-x]; ok {
			s := 0
			for i := k; i >= 0; i-- {
				s += c[k-i]
				ans += b[i] * s
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf525E(os.Stdin, os.Stdout) }
