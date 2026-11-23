package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1310B(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	hb := 1 << (n - 1)
	f := make([][]int, hb*2)
	for i := range f {
		f[i] = []int{0, -1e9, -1e9, -1e9}
	}
	for range k {
		var v int
		Fscan(in, &v)
		v = hb + (v-1)/2
		f[v][3] = f[v][1]
		f[v][1] = 1
		f[v][2] = 1
	}
	for i := hb - 1; i > 0; i-- {
		for x := range 4 {
			for y := range 4 {
				f[i][x|y] = max(f[i][x|y], f[i*2][x]+f[i*2+1][y]+(x|y))
			}
		}
	}
	Fprint(out, max(f[1][1]+1, f[1][2]+1, f[1][3]+1, 0))
}

//func main() { cf1310B(bufio.NewReader(os.Stdin), os.Stdout) }
