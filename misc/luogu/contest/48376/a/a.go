package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	for T := r(); T > 0; T-- {
		n, win, lose, tie, ga, gb := r(), r(), r(), r(), r(), r()
		mx, mi := tie, tie
		// 输一球
		if gb > 0 {
			mx = lose
		}
		winN := min(n-1, ga)
		mx += winN*win + (n-1-winN)*tie
		if ga >= gb { // 也可以不输球
			winN := min(n, ga-gb)
			mx = max(mx, winN*win+(n-winN)*tie)
		}

		// 赢一球
		if ga > 0 {
			mi = win
		}
		loseN := min(n-1, gb)
		mi += loseN*lose + (n-1-loseN)*tie
		if gb >= ga { // 也可以不赢球
			loseN := min(n, gb-ga)
			mi = min(mi, loseN*lose+(n-loseN)*tie)
		}
		Fprintln(out, mx, mi)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
