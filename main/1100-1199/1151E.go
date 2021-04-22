package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1151E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, pre, ans int64
	Fscan(in, &n)
	for m := n; m > 0; m-- {
		Fscan(in, &v)
		if v > pre {
			ans += (v - pre) * (n - v + 1)
		} else {
			ans += (pre - v) * v
		}
		pre = v
	}
	Fprint(out, ans)
}

//func main() { CF1151E(os.Stdin, os.Stdout) }
