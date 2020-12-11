package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1355B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int, n+2)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			cnt[v]++
		}
		ans := 0
		for i := 1; i <= n; i++ {
			ans += cnt[i] / i
			cnt[i+1] += cnt[i] % i
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1355B(os.Stdin, os.Stdout) }
