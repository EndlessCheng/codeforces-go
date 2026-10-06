package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p5329(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	var s string
	Fscan(in, &n, &s)

	ans := make([]int, n)
	l, r := 0, n-1
	st := 1
	for i := 1; i <= n; i++ {
		if i == n || s[i] > s[i-1] {
			for j := i; j >= st; j-- {
				ans[r] = j
				r--
			}
			st = i + 1
		} else if s[i] < s[i-1] {
			for j := st; j <= i; j++ {
				ans[l] = j
				l++
			}
			st = i + 1
		}
	}

	for _, i := range ans {
		Fprint(out, i, " ")
	}
}

//func main() { p5329(bufio.NewReader(os.Stdin), os.Stdout) }
