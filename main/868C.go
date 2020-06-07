package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF868C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, a int
	Fscan(in, &n, &k)
	has := make([]bool, 1<<k)
	for ; n > 0; n-- {
		v := 0
		for i := 0; i < k; i++ {
			Fscan(in, &a)
			v |= a << i
		}
		has[v] = true
	}
	if has[0] {
		Fprint(out, "YES")
		return
	}
	for i, v := range has {
		for j, w := range has {
			if v && w && j != i && i&j == 0 {
				Fprint(out, "YES")
				return
			}
		}
	}
	Fprint(out, "NO")
}

//func main() { CF868C(os.Stdin, os.Stdout) }
