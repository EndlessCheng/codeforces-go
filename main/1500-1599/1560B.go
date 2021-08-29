package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1560B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, a, b, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		m := abs(a - b)
		n := m * 2
		if a > n || b > n || c > n {
			Fprintln(out, -1)
		} else {
			Fprintln(out, (c-1+m)%n+1)
		}
	}
}

//func main() { CF1560B(os.Stdin, os.Stdout) }
