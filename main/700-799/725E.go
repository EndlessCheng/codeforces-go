package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf725E(in io.Reader, out io.Writer) {
	var c, n int
	Fscan(in, &c, &n)
	cnt := make([]int, c+1)
	pre := make([]int, c+1)
	for range n {
		var x int
		Fscan(in, &x)
		cnt[x]++
	}

	pre[0] = -1
	for i := 1; i <= c; i++ {
		if cnt[i] > 0 {
			pre[i] = i
		} else {
			pre[i] = pre[i-1]
		}
	}

	for i := 1; i <= c; i++ {
		x := c
		y := pre[c]
		ok := true
		for x > 0 && y != -1 {
			if y < i && x >= i && ok {
				ok = false
				x -= i
			} else {
				x -= min(cnt[y], x/y) * y
				y = pre[min(y-1, x)]
			}
		}
		if x != 0 {
			Fprint(out, i)
			return
		}
	}
	Fprint(out, "Greed is good")
}

//func main() { cf725E(bufio.NewReader(os.Stdin), os.Stdout) }
