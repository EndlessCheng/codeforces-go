package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1006F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	var k, ans int64
	Fscan(in, &n, &m, &k)
	t := (n + m - 2) / 2
	cnt := [20][20]map[int64]int{}
	a := make([][]int64, n)
	for i := range a {
		a[i] = make([]int64, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			if i+j == t {
				cnt[i][j] = map[int64]int{}
			}
		}
	}
	var f func(int, int, int64)
	f = func(x, y int, s int64) {
		s ^= a[x][y]
		if x+y == t {
			cnt[x][y][s]++
			return
		}
		if x+1 < n {
			f(x+1, y, s)
		}
		if y+1 < m {
			f(x, y+1, s)
		}
	}
	f(0, 0, 0)
	f = func(x, y int, s int64) {
		if x+y == t {
			ans += int64(cnt[x][y][s])
			return
		}
		s ^= a[x][y]
		if x > 0 {
			f(x-1, y, s)
		}
		if y > 0 {
			f(x, y-1, s)
		}
	}
	f(n-1, m-1, k)
	Fprint(out, ans)
}

//func main() { CF1006F(os.Stdin, os.Stdout) }
