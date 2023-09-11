package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1707A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := bytes.Repeat([]byte{'0'}, n)
		iq := 0
		for i := n - 1; i >= 0; i-- {
			if iq >= a[i] {
				ans[i] = '1'
			} else if iq < q {
				ans[i] = '1'
				iq++
			}
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { CF1707A(os.Stdin, os.Stdout) }
