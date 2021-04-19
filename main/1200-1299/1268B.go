package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1268B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	n, v, c := 0, int64(0), [2]int64{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		c[0] += v / 2
		c[1] += v / 2
		if v&1 > 0 {
			c[n&1]++
		}
	}
	if c[0] > c[1] {
		c[0] = c[1]
	}
	Fprint(out, c[0])
}

//func main() { CF1268B(os.Stdin, os.Stdout) }
