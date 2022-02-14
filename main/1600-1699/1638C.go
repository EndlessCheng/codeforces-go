package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1638C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := int64(0)
		for i, s := int64(1), int64(0); i <= n; i++ {
			Fscan(in, &v)
			s += v
			if s == i*(i+1)/2 {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1638C(os.Stdin, os.Stdout) }
