package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func p3188(in io.Reader, out io.Writer) {
	var n, maxW int
	for {
		Fscan(in, &n, &maxW)
		if n < 0 {
			break
		}

		width := bits.Len(uint(maxW))
		const mx = 1000
		f := make([][mx + 1]int, width)
		sumW := make([]int, width)
		for range n {
			var w, v int
			Fscan(in, &w, &v)
			tz := bits.TrailingZeros(uint(w))
			if tz >= width {
				continue
			}
			w >>= tz
			sumW[tz] += w
			for j := sumW[tz]; j >= w; j-- {
				f[tz][j] = max(f[tz][j], f[tz][j-w]+v)
			}
		}

		for i, s := range sumW {
			for j := 1; j <= s; j++ {
				f[i][j] = max(f[i][j], f[i][j-1])
			}
			if i+1 < width {
				sumW[i+1] = max(sumW[i+1], s/2)
			}
		}

		for i := 1; i < width; i++ {
			for j := sumW[i]; j >= 0; j-- {
				for k := range j + 1 {
					mx1 := min(k<<1|maxW>>(i-1)&1, sumW[i-1])
					f[i][j] = max(f[i][j], f[i][j-k]+f[i-1][mx1])
				}
			}
		}

		Fprintln(out, f[width-1][min(1, sumW[width-1])])
	}
}

//func main() { p3188(bufio.NewReader(os.Stdin), os.Stdout) }
