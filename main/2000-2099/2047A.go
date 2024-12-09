package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf2047A(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		ans := 0
		s := 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			s += v
			rt := int(math.Sqrt(float64(s)))
			if rt%2 > 0 && rt*rt == s {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2047A(bufio.NewReader(os.Stdin), os.Stdout) }
