package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1208C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	ans := make([][]interface{}, n)
	for i := range ans {
		ans[i] = make([]interface{}, n)
	}
	n /= 2
	for i, c := 0, 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans[i][j] = 4 * c
			ans[i][j+n] = 4*c + 1
			ans[i+n][j] = 4*c + 2
			ans[i+n][j+n] = 4*c + 3
			c++
		}
	}
	for _, row := range ans {
		Fprintln(out, row...)
	}
}

//func main() { CF1208C(os.Stdin, os.Stdout) }
