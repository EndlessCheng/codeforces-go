package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, v int
	var s string
	Fscan(in, &n, &s)
	f := [2][2]int{}
	for i, b := range s {
		Fscan(in, &v)
		f = [2][2]int{
			{f[0][1], f[0][0]},
			{min(f[1][1], f[0][0]), min(f[1][0], f[0][1])},
		}
		f[0][b-'0'^1] += v
		if i > 0 {
			f[1][b-'0'^1] += v
		} else {
			f[1] = [2]int{1e18, 1e18}
		}
	}
	Fprint(out, min(f[1][0], f[1][1]))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
