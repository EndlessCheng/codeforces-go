package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1508B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var k int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		if n < 61 && k > 1<<(n-1) {
			Fprintln(out, -1)
			continue
		}
		k--
		cur := 1
		for i := n - 2; i >= 0; i-- {
			if i > 59 || k>>i&1 == 0 {
				Fprint(out, cur, " ")
				cur++
				continue
			}
			st := i
			for ; i >= 0 && k>>i&1 > 0; i-- {
			}
			for j := st; j >= i; j-- {
				Fprint(out, cur+j-i, " ")
			}
			cur += st - i + 1
		}
		if cur == n {
			Fprint(out, n)
		}
		Fprintln(out)
	}
}

//func main() { CF1508B(os.Stdin, os.Stdout) }
