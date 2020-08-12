package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1097B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := func(sub int) (res int) {
		for i, v := range a {
			if sub>>i&1 == 0 {
				res += v
			} else {
				res -= v
			}
		}
		return
	}
	for sub := 0; sub < 1<<n; sub++ {
		if f(sub)%360 == 0 {
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF1097B(os.Stdin, os.Stdout) }
