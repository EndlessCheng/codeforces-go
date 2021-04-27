package main

import (
	. "fmt"
	"strings"
)

type (
	input23 struct{ n int }
	req23   struct{ sx, sy, tx, ty int }
	resp23  struct{ s string }
	guess23 struct{ ans string }
)

// github.com/EndlessCheng/codeforces-go
func CF1023E(in input23, Q func(req23) resp23) (gs guess23) {
	q := func(sx, sy, tx, ty int) bool { return Q(req23{sx, sy, tx, ty}).s[0] == 'Y' }
	n := in.n
	ans, t := "", ""
	defer func() { gs.ans = ans }()

	x0, y0 := 1, 1
	for x0+y0 <= n {
		if q(x0+1, y0, n, n) {
			x0++
			ans += "D"
		} else {
			y0++
			ans += "R"
		}
	}
	x, y := n, n
	for y > y0 {
		if q(1, 1, x, y-1) {
			y--
			t += "R"
		} else {
			x--
			t += "D"
		}
	}
	ans += strings.Repeat("D", x-x0)
	for i := len(t) - 1; i >= 0; i-- {
		ans += string(t[i])
	}
	return
}

func ioq23() {
	// Interaction
	Q := func(req req23) (resp resp23) {
		Println("?", req.sx, req.sy, req.tx, req.ty)
		Scan(&resp.s)
		return
	}

	// Input
	d := input23{}
	Scan(&d.n)

	// Output
	gs := CF1023E(d, Q)
	Println("!", gs.ans)
}

//func main() { ioq23() }
