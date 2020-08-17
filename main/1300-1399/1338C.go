package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1338C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	mp := [3][4]int64{{0, 1, 2, 3}, {0, 2, 3, 1}, {0, 3, 1, 2}}

	var t, n int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		n--
		tp := n % 3
		i := 0
		for n /= 3; 1<<i <= n; i += 2 {
			n -= 1 << i
		}
		tail := int64(0)
		for i := 0; n > 0; n >>= 2 {
			tail |= mp[tp][n&3] << i
			i += 2
		}
		Fprintln(out, int64(1)<<i*(tp+1)|tail)
	}
}

//func main() { CF1338C(os.Stdin, os.Stdout) }
