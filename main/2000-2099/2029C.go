package main

import (
	"bufio"
	. "cmp"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2029C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f := [3]int{0, -1e9, -1e9}
		for range n {
			Fscan(in, &v)
			f[2] = max(f[2]+Compare(v, f[2]), f[1]+Compare(v, f[1]))
			f[1] = max(f[1], f[0])
			f[0] += Compare(v, f[0])
		}
		Fprintln(out, max(f[1], f[2]))
	}
}

//func main() { cf2029C(bufio.NewReader(os.Stdin), os.Stdout) }
