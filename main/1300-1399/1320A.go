package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1320A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int64
	Fscan(in, &n)
	c := map[int64]int64{}
	for i := int64(0); i < n; i++ {
		Fscan(in, &v)
		c[v-i] += v
		if c[v-i] > ans {
			ans = c[v-i]
		}
	}
	Fprint(out, ans)
}

//func main() { CF1320A(os.Stdin, os.Stdout) }
