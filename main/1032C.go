package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1032C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 1)
		return
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	if a[0] < a[1] {
		b[0] = 1
	} else {
		b[0] = 5
	}
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			if i > 1 && a[i-1] <= a[i-2] {
				b[i-1] = 1
				if b[i-2] == 1 {
					b[i-1] = 2
				}
			}
			b[i] = b[i-1] + 1
		} else if a[i] < a[i-1] {
			if i > 1 && a[i-1] >= a[i-2] {
				b[i-1] = 5
				if b[i-2] == 5 {
					b[i-1] = 4
				}
			}
			b[i] = b[i-1] - 1
		} else {
			b[i] = (b[i-1]-1)%3 + 2
		}
		if b[i] < 1 || b[i] > 5 {
			Fprint(out, -1)
			return
		}
	}
	for _, v := range b {
		Fprint(out, v, " ")
	}
}

//func main() { CF1032C(os.Stdin, os.Stdout) }
