package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1328D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
o:
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		same := true
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if i > 0 && a[i] != a[i-1] {
				same = false
			}
		}
		if same {
			Fprintln(out, 1)
			Fprintln(out, strings.Repeat("1 ", n))
			continue o
		}
		if n&1 == 0 {
			Fprintln(out, 2)
			Fprintln(out, strings.Repeat("1 2 ", n/2))
			continue o
		}
		for i, v := range a {
			if v == a[(i+1)%n] {
				ans := make([]interface{}, n)
				for j, c := i, 1; j >= 0; j-- {
					ans[j] = c
					c = 3 - c
				}
				for j, c := i+1, 1; j < n; j++ {
					ans[j] = c
					c = 3 - c
				}
				Fprintln(out, 2)
				Fprintln(out, ans...)
				continue o
			}
		}
		Fprintln(out, 3)
		Fprintln(out, strings.Repeat("1 2 ", n/2)+"3")
	}
}

//func main() { CF1328D(os.Stdin, os.Stdout) }
