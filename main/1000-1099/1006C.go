package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1006C(in io.Reader, out io.Writer) {
	var n, ans, pre, suf int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	r := n - 1
	for i, v := range a {
		pre += v
		for ; r >= 0 && suf < pre; r-- {
			suf += a[r]
		}
		if r < i {
			break
		}
		if suf == pre {
			ans = max(ans, pre)
		}
	}
	Fprint(out, ans)
}

//func main() { cf1006C(bufio.NewReader(os.Stdin), os.Stdout) }
