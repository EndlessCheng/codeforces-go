package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1154G(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	id := [1e7 + 1]int{}
	min := int64(1e18)
	var n, v, x, y int
	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if id[v] == 0 {
			id[v] = i
		} else if int64(v) < min {
			min, x, y = int64(v), id[v], i
		}
	}
	for i := 1; i <= 1e7; i++ {
		v = 0
		for j := i; j <= 1e7; j += i {
			if id[j] > 0 {
				if v > 0 {
					if l := int64(v) * int64(j) / int64(i); l < min {
						min, x, y = l, id[v], id[j]
					}
					break
				}
				v = j
			}
		}
	}
	if x > y {
		x, y = y, x
	}
	Fprint(out, x, y)
}

//func main() { CF1154G(os.Stdin, os.Stdout) }
