package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1490C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	c := map[int64]bool{}
	for i := int64(1); i*i*i < 1e12; i++ {
		c[i*i*i] = true
	}

	var T int
	var x int64
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		for v := range c {
			if c[x-v] {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1490C(os.Stdin, os.Stdout) }
