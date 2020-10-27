package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1175A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ans := int64(-1)
		for n > 0 {
			ans += n%k + 1
			n /= k
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1175A(os.Stdin, os.Stdout) }
