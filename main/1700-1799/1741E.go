package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1741E(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f := make([]bool, n+1)
		f[0] = true
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if v < i && f[i-1-v] {
				f[i] = true
			}
			if i+v <= n && f[i-1] {
				f[i+v] = true
			}
		}
		if f[n] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1741E(bufio.NewReader(os.Stdin), os.Stdout) }
