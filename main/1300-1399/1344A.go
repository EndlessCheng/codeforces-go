package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1344A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		has := make([]bool, n)
		for i, v := range a {
			v = ((v+i)%n + n) % n
			if has[v] {
				Fprintln(out, "NO")
				continue o
			}
			has[v] = true
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1344A(os.Stdin, os.Stdout) }
