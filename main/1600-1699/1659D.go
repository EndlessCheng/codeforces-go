package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1659D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c := make([]int, n)
		s := int64(0)
		for i := range c {
			Fscan(in, &c[i])
			s += int64(c[i])
		}
		k := int(s / int64(n))

		d := make([]int, n)
		sd := 0
		a := make([]int, n)
		for i := n - 1; i >= 0; i-- {
			sd += d[i]
			if c[i]+sd == i+1 {
				a[i] = 1
			}
			sd--
			if i-k >= 0 {
				d[i-k]++
			}
			if a[i] == 1 {
				k--
			}
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1659D(os.Stdin, os.Stdout) }
