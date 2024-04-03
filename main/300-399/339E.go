package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf339E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n int
	Fscan(in, &n)
	a := make([]int, n+2)
	a[0] = -1
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	a[n+1] = -1

	ans := [3][2]int{}
	var f func(int) bool
	f = func(k int) bool {
		for i := 1; i <= n; i++ {
			if a[i] != i {
				goto next
			}
		}
		Fprintln(out, 2-k)
		for _, p := range ans[k+1:] {
			Fprintln(out, p[0], p[1])
		}
		return true
	next:
		if k < 0 {
			return false
		}
		for i := 1; i <= n; i++ {
			if a[i] == i || abs(a[i]-a[i-1]) == 1 && abs(a[i]-a[i+1]) == 1 {
				continue
			}
			for j := i + 1; j <= n; j++ {
				if a[j] == j || abs(a[j]-a[j-1]) == 1 && abs(a[j]-a[j+1]) == 1 {
					continue
				}
				slices.Reverse(a[i : j+1])
				ans[k] = [2]int{i, j}
				if f(k - 1) {
					return true
				}
				slices.Reverse(a[i : j+1])
			}
		}
		return false
	}
	f(2)
}

//func main() { cf339E(os.Stdin, os.Stdout) }
