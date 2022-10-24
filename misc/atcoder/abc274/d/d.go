package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n, &x, &y)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	x -= a[0]
	a = a[1:]
	f := func(a []int, x int) bool {
		tot := x
		for _, v := range a {
			tot += v
		}
		if tot&1 > 0 || tot < 0 {
			return false
		}
		tot /= 2
		dp := make([]bool, tot+1)
		dp[0] = true
		for _, v := range a {
			for j := tot; j >= v; j-- {
				dp[j] = dp[j] || dp[j-v]
			}
		}
		return dp[tot]
	}
	var b, c []int
	for i, v := range a {
		if i&1 == 0 {
			b = append(b, v)
		} else {
			c = append(c, v)
		}
	}
	if f(c, x) && f(b, y) {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}
}

func main() { run(os.Stdin, os.Stdout) }
