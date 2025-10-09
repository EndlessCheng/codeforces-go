package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2140C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s, mx := 0, n-2+n%2
		minPos, maxNeg := int(1e18), int(-1e18)
		for i := range n {
			Fscan(in, &v)
			if i%2 > 0 {
				s -= v
				mx = max(mx, i+v*2-minPos)
				maxNeg = max(maxNeg, v*2-i)
			} else {
				s += v
				mx = max(mx, i-v*2+maxNeg)
				minPos = min(minPos, v*2+i)
			}
		}
		Fprintln(out, s+mx)
	}
}

//func main() { cf2140C(bufio.NewReader(os.Stdin), os.Stdout) }
