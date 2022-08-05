package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF280B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, v, ans int
	st := []int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		for len(st) > 0 && st[len(st)-1] < v {
			ans = max(ans, st[len(st)-1]^v)
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans = max(ans, st[len(st)-1]^v)
		}
		st = append(st, v)
	}
	Fprint(out, ans)
}

//func main() { CF280B(os.Stdin, os.Stdout) }
