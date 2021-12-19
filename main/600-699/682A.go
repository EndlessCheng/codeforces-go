package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF682A(in io.Reader, out io.Writer) {
	var n, m, ans int64
	Fscan(in, &n, &m)
	for i := int64(1); i <= 5; i++ {
		j := (5-i)%5 + i - i%5
		ans += (n - i + 5) / 5 * ((m - j + 5) / 5)
	}
	Fprint(out, ans)
}

//func main() { CF682A(os.Stdin, os.Stdout) }
