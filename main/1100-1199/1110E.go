package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1110E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, p, v int
	Fscan(in, &n, &p)
	p0 := p
	c := make(map[int]int, n-1)
	for i := 1; i < n; i++ {
		Fscan(in, &v)
		c[v-p]++
		p = v
	}
	Fscan(in, &p)
	if p != p0 {
		Fprint(out, "No")
		return
	}
	for n--; n > 0; n-- {
		Fscan(in, &v)
		c[v-p]--
		p = v
	}
	for _, v := range c {
		if v != 0 {
			Fprint(out, "No")
			return
		}
	}
	Fprint(out, "Yes")
}

//func main() { CF1110E(os.Stdin, os.Stdout) }
