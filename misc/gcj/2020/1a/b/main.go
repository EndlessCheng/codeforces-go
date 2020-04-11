package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(_case int) {
		var n int
		Fscan(in, &n)
		if n <= 500 {
			for i := 1; i <= n; i++ {
				Fprintln(out, i, 1)
			}
			return
		}
		m := n - 31
		sum := 0
		left := true
		for i := uint(0); sum < n; i++ {
			if m>>i&1 == 1 {
				if left {
					for k := uint(1); k <= i+1; k++ {
						Fprintln(out, i+1, k)
					}
				} else {
					for k := i + 1; k > 0; k-- {
						Fprintln(out, i+1, k)
					}
				}
				left = !left
				sum += 1 << i
			} else {
				if left {
					Fprintln(out, i+1, 1)
				} else {
					Fprintln(out, i+1, i+1)
				}
				sum++
			}
		}
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d:\n", _case)
		solve(_case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
