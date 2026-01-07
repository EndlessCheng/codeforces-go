package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf316D3(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, v, c1 int
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		c1 += 2 - v
	}

	f := make([]int, c1+2)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= c1; i++ {
		f[i] = (f[i-1] + (i-1)*f[i-2]) % mod
	}

	ans := f[c1]
	for i := c1 + 1; i <= n; i++ {
		ans = ans * i % mod
	}
	Fprint(out, ans)
}

//func main() { cf316D3(bufio.NewReader(os.Stdin), os.Stdout) }
