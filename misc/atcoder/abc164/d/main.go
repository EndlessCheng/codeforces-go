package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	const m = 2019
	var exgcd func(a, b int) (x, y int)
	exgcd = func(a, b int) (x, y int) {
		if b == 0 {
			return 1, 0
		}
		y, x = exgcd(b, a%b)
		y -= a / b * x
		return
	}
	invM := func(a int) int { x, _ := exgcd(a, m); return (x%m + m) % m }

	var s []byte
	Fscan(bufio.NewReader(_r), &s)
	cnts := [m]int{0: 1}
	ans, v, ten := 0, 0, 1
	for _, b := range s {
		v = (v*10 + int(b-'0')) % m
		w := v * invM(ten) % m
		ans += cnts[w]
		cnts[w]++
		ten = ten * 10 % m
	}
	Fprint(_w, ans)
}

func main() { run(os.Stdin, os.Stdout) }
