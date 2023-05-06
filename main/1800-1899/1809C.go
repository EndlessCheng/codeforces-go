package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1809C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ans := make([]int, 0, n)
		for k > len(ans) {
			ans = append(ans, 2)
			k -= len(ans)
		}
		if k > 0 {
			ans = append(ans, (k-len(ans))*2-1)
		}
		for len(ans) < n {
			ans = append(ans, -1000)
		}
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1809C(os.Stdin, os.Stdout) }
