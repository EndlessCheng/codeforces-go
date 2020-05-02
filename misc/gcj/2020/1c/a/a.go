package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	dir4 := [...][2]int{
		'N': {0, 1},
		'S': {0, -1},
		'E': {1, 0},
		'W': {-1, 0},
	}

	solve := func(_case int) {
		var x, y int
		var s string
		Fscan(in, &x, &y, &s)
		s = "A" + s
		for i, b := range s {
			d := dir4[b]
			x += d[0]
			y += d[1]
			if abs(x)+abs(y) <= i {
				Fprintln(out, i)
				return
			}
		}
		Fprintln(out, "IMPOSSIBLE")
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: ", _case)
		solve(_case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
