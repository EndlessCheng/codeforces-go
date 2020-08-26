package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1400D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := int64(0)
		cnt := make([]int64, n+1)
		cnt[a[0]]++
		for j := 1; j < n-2; j++ {
			v := a[j]
			c := cnt[a[j+1]]
			for _, w := range a[j+2:] {
				if w == v {
					ans += c
				}
				c += cnt[w]
			}
			cnt[v]++
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1400D(os.Stdin, os.Stdout) }
