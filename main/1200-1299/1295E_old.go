package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1295E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	p := make([]int, n)
	pos := make([]int, n)
	for i := range p {
		Fscan(in, &p[i])
		p[i]--
		pos[p[i]] = i
	}
	if p[0] == 0 || p[n-1] == n-1 {
		Fprint(out, 0)
		return
	}
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := min(a[0], a[n-1])
	moveToL := make([]bool, n)
	moveToR := make([]bool, n)
	moveToL[0] = true
	moveToR[p[0]] = true
	cost := a[0] + a[pos[0]]
	for i := 1; i < n-1; i++ {
		if pi := p[i]; pi > i {
			if !moveToR[pi] {
				moveToR[pi] = true
				cost += a[i]
			}
			if moveToL[pi] {
				moveToL[pi] = false
				cost -= a[i]
			}
		} else {
			if moveToL[pi] {
				moveToL[pi] = false
				cost -= a[i]
			}
		}
		if moveToR[i] {
			moveToR[i] = false
			cost -= a[pos[i]]
		}
		if !moveToL[i] && pos[i] > i { // else if
			moveToL[i] = true
			cost += a[pos[i]]
		}
		ans = min(ans, cost)
	}
	Fprint(out, ans)
}

//func main() {	CF1295E(os.Stdin, os.Stdout) }
