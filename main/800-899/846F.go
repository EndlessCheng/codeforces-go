package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf846F(in io.Reader, out io.Writer) {
	var n, v, ans int
	Fscan(in, &n)
	pre := [1e6 + 1]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		ans += (i - pre[v]) * (n - i + 1)
		pre[v] = i
	}
	Fprintf(out, "%.4f", float64(ans*2-n)/float64(n*n))
}

//func main() { cf846F(bufio.NewReader(os.Stdin), os.Stdout) }
