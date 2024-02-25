package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf650A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y int }
	cx := map[int]int{}
	cy := map[int]int{}
	cp := map[pair]int{}
	var n, x, y, ans int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x, &y)
		p := pair{x, y}
		ans += cx[x] + cy[y] - cp[p]
		cx[x]++
		cy[y]++
		cp[p]++
	}
	Fprint(out, ans)
}

//func main() { cf650A(os.Stdin, os.Stdout) }
