package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1334B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, x int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &x)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		sum := int64(0)
		cnt := int64(0)
		ans := 0
		for i := n - 1; i >= 0; i-- {
			if v := a[i]; v > x {
				sum += int64(v)
				cnt++
				ans = n - i
			} else if v == x {
				ans = n - i
			} else {
				sum += int64(v)
				cnt++
				if int64(x)*cnt <= sum {
					ans = n - i
				} else {
					break
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1334B(os.Stdin, os.Stdout) }
