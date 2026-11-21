package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func cf2254C2(in io.Reader, out io.Writer) {
	var T, n int
	var s, t string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		pos := [2][]int{}
		for i, b := range s {
			if b == '1' {
				pos[i&1] = append(pos[i&1], i)
			}
		}

		ans := 0
		for i, b := range t {
			if b == '1' {
				p := i & 1
				if len(pos[p]) == 0 {
					Fprintln(out, -1)
					continue o
				}
				ans += abs(pos[p][0] - i)
				pos[p] = pos[p][1:]
			}
		}
		if len(pos[0]) > 0 || len(pos[1]) > 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, ans/2)
		}
	}
}

func main() { cf2254C2(bufio.NewReader(os.Stdin), os.Stdout) }
func abs(x int) int { if x < 0 { return -x }; return x }
