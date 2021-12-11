package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF467B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v, ans uint
	Fscan(in, &n, &m, &k)
	a := make([]uint, m)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &v)

	for _, w := range a {
		if uint(bits.OnesCount(v^w)) <= k {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { CF467B(os.Stdin, os.Stdout) }
