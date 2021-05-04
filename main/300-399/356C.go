package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF356C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	c := [5]int{}
	var n, v, s int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		c[v]++
		s += v
	}
	if s < 3 || s == 5 {
		Fprint(out, -1)
		return
	}
	mi := c[1]
	if c[2] < mi {
		mi = c[2]
	}
	c[1] -= mi
	c[2] -= mi
	c[3] += mi
	ans := mi + (c[1]+c[2])/3*2
	if c[1] > 0 {
		c[3] += c[1] / 3
	} else {
		c[3] += c[2] / 3 * 2
	}
	if c[1] %= 3; c[1] == 1 && c[3] == 0 {
		ans += 2
	} else {
		ans += c[1]
	}
	if c[2] %= 3; c[2] == 1 && c[4] == 0 {
		ans += 2
	} else {
		ans += c[2]
	}
	Fprint(out, ans)
}

//func main() { CF356C(os.Stdin, os.Stdout) }
