package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2044D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		vis := make([]bool, n+1)
		for range n {
			Fscan(in, &v)
			if !vis[v] {
				vis[v] = true
				Fprint(out, v, " ")
			}
		}
		for i := 1; i <= n; i++ {
			if !vis[i] {
				Fprint(out, i, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2044D(bufio.NewReader(os.Stdin), os.Stdout) }
