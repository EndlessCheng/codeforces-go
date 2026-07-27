package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2242D(in io.Reader, out io.Writer) {
	var T int
	for Fscan(in, &T); T > 0; T-- {
		calcSum := func() []int {
			var s string
			Fscan(in, &s)
			sum := make([]int, len(s)+1)
			for i, b := range s {
				sum[i+1] = (sum[i] + int(b-'0')) % 10
			}
			return sum
		}
		a := calcSum()[1:]
		b := calcSum()[1:]
		if a[len(a)-1] != b[len(b)-1] {
			Fprintln(out, -1)
			continue
		}

		m := len(b)
		f := make([]int, m+1)
		for _, x := range a {
			pre := 0
			for j, y := range b {
				if x == y {
					f[j+1], pre = pre+1, f[j+1]
				} else {
					pre = f[j+1]
					f[j+1] = max(f[j+1], f[j])
				}
			}
		}
		Fprintln(out, f[m])
	}
}

//func main() { cf2242D(bufio.NewReader(os.Stdin), os.Stdout) }
