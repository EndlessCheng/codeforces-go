package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1017C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, p int
	Fscan(in, &n)
	min := n * 2
	for i := 1; i <= n; i++ {
		if res := i + (n-1)/i; res < min { // 这里 res 省略了 +1
			min, p = res, i
		}
	}
	// p 可以直接取 sqrt(n)
	for i, cur := 0, n; i < p; i++ {
		sz := n / p
		if i < n%p {
			sz++
		}
		for j := cur - sz + 1; j <= cur; j++ {
			Fprint(out, j, " ")
		}
		cur -= sz
	}
}

//func main() { CF1017C(os.Stdin, os.Stdout) }
