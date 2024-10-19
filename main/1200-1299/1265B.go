package main

import (
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1265B(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([]int, n)
		for i := range n {
			Fscan(in, &v)
			pos[v-1] = i
		}
		mn, mx := n, -1
		ans := bytes.Repeat([]byte{'0'}, n)
		for i, p := range pos {
			mn = min(mn, p)
			mx = max(mx, p)
			if mx-mn == i {
				ans[i] = '1'
			}
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { cf1265B(bufio.NewReader(os.Stdin), os.Stdout) }
