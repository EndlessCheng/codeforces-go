package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1054B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, mex int
	ans := -1
	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if v > mex {
			ans = i
			break
		}
		if v == mex {
			mex++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1054B(os.Stdin, os.Stdout) }
