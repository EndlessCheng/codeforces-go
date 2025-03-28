package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2758(in io.Reader, out io.Writer) {
	var s, t []byte
	Fscan(in, &s, &t)
	m := len(t)
	f := make([]int, m+1)
	for j := 1; j <= m; j++ {
		f[j] = j
	}
	for _, x := range s {
		pre := f[0]
		f[0]++
		for j, y := range t {
			if x == y {
				f[j+1], pre = pre, f[j+1]
			} else {
				f[j+1], pre = min(f[j+1], f[j], pre)+1, f[j+1]
			}
		}
	}
	Fprint(out, f[m])
}

//func main() { p2758(bufio.NewReader(os.Stdin), os.Stdout) }
