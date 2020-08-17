package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1326C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v int
	Fscan(in, &n, &k)
	pos := make([]int, 0, k)
	for i := 0; i < n; i++ {
		if Fscan(in, &v); v > n-k {
			pos = append(pos, i)
		}
	}
	ans := int64(1)
	for i := 1; i < k; i++ {
		ans = ans * int64(pos[i]-pos[i-1]) % 998244353
	}
	Fprintln(out, int64(2*n-k+1)*int64(k)/2, ans)
}

//func main() { CF1326C(os.Stdin, os.Stdout) }
