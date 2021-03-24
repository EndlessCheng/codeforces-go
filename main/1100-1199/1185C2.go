package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1185C2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	cnt := [101]int{}
	var n, m, t, sum int
	for Fscan(in, &n, &m); n > 0; n-- {
		Fscan(in, &t)
		s, c := sum+t-m, 0
		for i := 100; s > 0 && i > 0; i-- { // 漏写了 s > 0 WA 了一次
			if s <= cnt[i]*i {
				c += (s + i - 1) / i
				break
			}
			c += cnt[i]
			s -= cnt[i] * i
		}
		Fprint(out, c, " ")
		sum += t
		cnt[t]++
	}
}

//func main() { CF1185C2(os.Stdin, os.Stdout) }
