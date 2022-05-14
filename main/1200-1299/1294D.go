package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1294D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, x, v, mex int
	Fscan(in, &q, &x)
	cnt := make([]int, x)
	for ; q > 0; q-- {
		Fscan(in, &v)
		cnt[v%x]++
		for cnt[mex%x] >= mex/x+1 {
			mex++
		}
		Fprintln(out, mex)
	}
}

//func main() { CF1294D(os.Stdin, os.Stdout) }
