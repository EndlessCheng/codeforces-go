package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
var a67 [1e6 + 1]int
var g67 [1e6 + 1][]int
var rt67, tot67, x67, y67 int

func f67(v int) int {
	s := a67[v]
	for _, w := range g67[v] {
		s += f67(w)
	}
	if v != rt67 && s == tot67 {
		if x67 == 0 {
			x67 = v
			s = 0
		} else if y67 == 0 {
			y67 = v
			s = 0
		}
	}
	return s
}

func CF767C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	for w := 1; w <= n; w++ {
		Fscan(in, &v, &a67[w])
		tot67 += a67[w]
		if v == 0 {
			rt67 = w
		} else {
			g67[v] = append(g67[v], w)
		}
	}
	if tot67%3 != 0 {
		Fprint(out, -1)
		return
	}
	tot67 /= 3

	f67(rt67)
	if y67 == 0 {
		Fprint(out, -1)
	} else {
		Fprint(out, x67, y67)
	}
}

//func main() { CF767C(os.Stdin, os.Stdout) }
