package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2192B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s string
	var ans []int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		pos := [2][]int{}
		for i, b := range s {
			pos[b-'0'] = append(pos[b-'0'], i)
		}

		if len(pos[1])&1 == 0 {
			ans = pos[1]
		} else if n&1 == 0 {
			ans = pos[0]
		} else {
			Fprintln(out, -1)
			continue
		}
		Fprintln(out, len(ans))
		for _, i := range ans {
			Fprint(out, i+1, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2192B(bufio.NewReader(os.Stdin), os.Stdout) }
