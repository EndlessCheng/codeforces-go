package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1282(in io.Reader, out io.Writer) {
	var n, x, y int
	Fscan(in, &n)
	bias := n * 5

	f := make([]int, bias*2+1)
	for i := range n {
		Fscan(in, &x, &y)
		nf := make([]int, len(f))
		for j := range nf {
			nf[j] = 1e9
		}
		d := x - y
		for j := bias - i*5; j <= bias+i*5; j++ {
			nf[j+d] = min(nf[j+d], f[j])
			nf[j-d] = min(nf[j-d], f[j]+1)
		}
		f = nf
	}

	for i := range bias + 1 {
		ans := min(f[bias-i], f[bias+i])
		if ans < 1e9 {
			Fprintln(out, ans)
			break
		}
	}
}

//func main() { p1282(bufio.NewReader(os.Stdin), os.Stdout) }
