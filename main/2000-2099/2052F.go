package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2052F(in io.Reader, out io.Writer) {
	ans := [3]string{"None", "Unique", "Multiple"}
	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		f := make([][3]int, n+1)
		f[0][0] = 1
		for i, v := range s {
			w := t[i]
			if v == '.' && w == '.' {
				f[i+1][0] = f[i][0]
				if i > 0 && s[i-1] == '.' && t[i-1] == '.' {
					f[i+1][0] = min(f[i+1][0]+f[i-1][0], 2)
				}
				f[i+1][1] = f[i][2]
				f[i+1][2] = f[i][1]
			} else if v == '.' && w == '#' {
				f[i+1][0] = f[i][1]
				f[i+1][1] = f[i][0]
			} else if v == '#' && w == '.' {
				f[i+1][0] = f[i][2]
				f[i+1][2] = f[i][0]
			} else {
				f[i+1][0] = f[i][0]
			}
		}
		Fprintln(out, ans[f[n][0]])
	}
}

//func main() { cf2052F(bufio.NewReader(os.Stdin), os.Stdout) }
