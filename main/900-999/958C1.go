package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf958C1(in io.Reader, out io.Writer) {
	var n, p, l, r, ans int
	Fscan(in, &n, &p)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		r += a[i]
	}
	for _, v := range a[:n-1] {
		r -= v
		l += v
		ans = max(ans, l%p+r%p)
	}
	Fprint(out, ans)
}

//func main() { cf958C1(bufio.NewReader(os.Stdin), os.Stdout) }
