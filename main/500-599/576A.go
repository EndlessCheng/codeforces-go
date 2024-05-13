package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf576A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	ans := []any{}
	notPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if notPrime[i] {
			continue
		}
		for v := i; v <= n; v *= i {
			ans = append(ans, v)
		}
		for j := i * i; j <= n; j += i {
			notPrime[j] = true
		}
	}
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

//func main() { cf576A(os.Stdin, os.Stdout) }
