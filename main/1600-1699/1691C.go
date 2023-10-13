package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF1691C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		i := strings.IndexByte(s, '1')
		if i < 0 {
			Fprintln(out, 0)
			continue
		}
		j := strings.LastIndexByte(s, '1')
		ans := strings.Count(s, "1") * 11
		if n-1-j <= k {
			k -= n - 1 - j
			ans -= 10
		}
		if ans > 1 && i <= k {
			ans--
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1691C(os.Stdin, os.Stdout) }
