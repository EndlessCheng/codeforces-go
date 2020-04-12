package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF553A(_r io.Reader, _w io.Writer) {
	const mod int64 = 1e9 + 7
	const mx int = 1e3
	C := [mx + 1][mx + 1]int64{}
	for i := 0; i <= mx; i++ {
		C[i][0] = 1
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % mod
		}
		C[i][i] = 1
	}

	in := bufio.NewReader(_r)
	var n, c, sum int
	ans := int64(1)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &c)
		sum += c
		ans = ans * C[sum-1][c-1] % mod
	}
	Fprint(_w, ans)
}

//func main() { CF553A(os.Stdin, os.Stdout) }
