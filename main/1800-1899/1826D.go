package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1826D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f := [3]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			f[2] = max(f[2], f[1]+v-i)
			f[1] = max(f[1], f[0]+v)
			f[0] = max(f[0], v+i)
		}
		Fprintln(out, f[2])
	}
}

//func main() { CF1826D(os.Stdin, os.Stdout) }
