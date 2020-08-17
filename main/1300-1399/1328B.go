package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1328B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		ans := bytes.Repeat([]byte{'a'}, n)
		i := 1
		for ; k > i; i++ {
			k -= i
		}
		ans[n-1-i] = 'b'
		ans[n-k] = 'b'
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { CF1328B(os.Stdin, os.Stdout) }
