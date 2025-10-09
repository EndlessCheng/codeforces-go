package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func f34(a, b int) (ans int) {
	for a > 1 {
		ans += b / a
		a, b = b%a, a
	}
	if a == 1 {
		return ans + b - 1
	}
	return 1e9
}

func cf134B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	ans := n
	for i := 1; i <= n; i++ {
		ans = min(ans, f34(i, n))
	}
	Fprint(out, ans)
}

//func main() { cf134B(os.Stdin, os.Stdout) }
