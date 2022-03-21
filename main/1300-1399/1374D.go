package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1374D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		cnt, mx := map[int]int{}, 0
		for Fscan(in, &n, &k); n > 0; n-- {
			Fscan(in, &v)
			if v = k - v%k; v < k {
				if cnt[v]++; cnt[v] > cnt[mx] || cnt[v] == cnt[mx] && v > mx {
					mx = v
				}
			}
		}
		if mx == 0 {
			Fprintln(out, 0)
		} else {
			Fprintln(out, int64(k)*int64(cnt[mx]-1)+int64(mx+1))
		}
	}
}

//func main() { CF1374D(os.Stdin, os.Stdout) }
