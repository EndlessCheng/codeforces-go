package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1800D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans := n - 1
		for i := 2; i < n; i++ {
			if s[i] == s[i-2] {
				ans--
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1800D(os.Stdin, os.Stdout) }
