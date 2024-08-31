package main

import (
	"bufio"
	. "fmt"
	"io"
)

func mexStr(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(s string) string {
		n := len(s)
		f := make([][2]int, n+1)
		for i := n - 1; i >= 0; i-- {
			f[i] = f[i+1]
			f[i][s[i]&1] = min(f[i+1][0], f[i+1][1]) + 1
		}
		if f[0][0] == 0 {
			return "0"
		}

		res := []byte{'1'}
		cur := f[0][1] + 1
		for i := 0; i < n-1; i++ {
			if res[len(res)-1] == s[i] {
				cur--
				if cur <= f[i+1][0] {
					res = append(res, '1')
				} else {
					res = append(res, '0')
				}
			}
		}
		for len(res) < f[0][1]+1 {
			res = append(res, '0')
		}
		return string(res)
	}

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		Fprintln(out, solve(s))
	}
}

//func main() { mexStr(bufio.NewReader(os.Stdin), os.Stdout) }
