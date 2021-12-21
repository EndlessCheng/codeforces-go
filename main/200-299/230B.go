package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF230B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e6
	np := [mx + 1]bool{1: true}
	for i := 2; i <= mx; i++ {
		if !np[i] {
			for j := 2 * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}

	var n, v int64
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		rt := int64(math.Sqrt(float64(v)))
		if rt*rt == v && !np[rt] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF230B(os.Stdin, os.Stdout) }
