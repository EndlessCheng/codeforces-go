package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2196D(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		t := [2]int{}
		q := [2][]int{make([]int, n+1), make([]int, n+1)}
		las := -1
		cnt := 0
		for i, b := range s {
			tp := 1
			if b == '(' || b == ')' {
				tp = 0
			}
			if b == '(' || b == '[' {
				t[tp]++
				q[tp][t[tp]] = i + 1
			} else if t[tp] > 0 {
				t[tp]--
			} else {
				las = i + 1
				cnt++
			}
		}

		x := n + 1
		if t[0] > 0 {
			x = q[0][1]
		}
		y := n + 1
		if t[1] > 0 {
			y = q[1][1]
		}

		ans := (cnt + t[0] + t[1]) / 2
		if cnt%2 == 1 && las < min(x, y) {
			ans++
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2196D(bufio.NewReader(os.Stdin), os.Stdout) }
