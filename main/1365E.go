package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1365E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int64
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	if n == 1 {
		ans = a[0]
	} else if n == 2 {
		ans = a[0] | a[1]
	} else {
		for i, v := range a {
			for j, w := range a[i+1:] {
				for _, u := range a[j+1:] {
					if v|w|u > ans {
						ans = v | w | u
					}
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1365E(os.Stdin, os.Stdout) }
