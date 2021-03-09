package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1493A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ans := []interface{}{}
		for i := (k + 1) / 2; i < k; i++ {
			ans = append(ans, i)
		}
		for i := k + 1; i <= n; i++ {
			ans = append(ans, i)
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { CF1493A(os.Stdin, os.Stdout) }
