package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1188B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	c := map[int]int{}
	var n int
	var p, k, v, ans int64
	for Fscan(in, &n, &p, &k); n > 0; n-- {
		Fscan(in, &v)
		w := int((v*v%p*v%p*v%p - v*k%p + p) % p)
		ans += int64(c[w])
		c[w]++
	}
	Fprint(out, ans)
}

//func main() { CF1188B(os.Stdin, os.Stdout) }
