package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF71C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]bool, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ps := []int{}
	x := n
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			for x /= i; x%i == 0; x /= i {
			}
			if i == 2 {
				if n%4 == 0 {
					ps = append(ps, 4)
				}
			} else {
				ps = append(ps, i)
			}
		}
	}
	if x > 1 {
		ps = append(ps, x)
	}

	for _, p := range ps {
		step := n / p
	o:
		for i := range a[:step] {
			for j := i; j < n; j += step {
				if !a[j] {
					continue o
				}
			}
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

func main() { CF71C(os.Stdin, os.Stdout) }
