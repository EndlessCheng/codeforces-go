package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1073D(in io.Reader, out io.Writer) {
	var n, t, ans int
	Fscan(in, &n, &t)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for {
		s, c := 0, 0
		for _, v := range a {
			if s+v <= t {
				s += v
				c++
			}
		}
		if c == 0 {
			Fprint(out, ans)
			return
		}
		ans += t / s * c
		t %= s
	}
}

//func main() { cf1073D(bufio.NewReader(os.Stdin), os.Stdout) }
