package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1389D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var T, n, k, l1, r1, l2, r2, ans int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &l1, &r1, &l2, &r2)
		if l1 > l2 {
			l1, r1, l2, r2 = l2, r2, l1, r1
		}
		d, D := min(r1, r2)-l2, max(r1, r2)-l1 // d=已经重合的长度，D=完全重合的长度
		if d > 0 {
			k -= n * d
			if k <= 0 {
				Fprintln(out, 0)
				continue
			}
			D -= d
			d = 0
		} else {
			d = -d
		}
		if k <= D {
			ans = d + k // 只需延长一对
		} else {
			ans = d + D + (k-D)*2 // 完全延长一对+在这对的基础上延长
			if k <= n*D {
				ans = min(ans, k/D*(d+D)+min(k%D*2, d+k%D)) // 完全延长 k/D 对，剩余的 k，要么在某对完全延长的区间上延长，要么用一对没有延长过的区间
			} else {
				ans = min(ans, n*(d+D)+(k-n*D)*2) // 完全延长 n 对，剩余的 k 在某对完全延长的区间上延长
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1389D(os.Stdin, os.Stdout) }
