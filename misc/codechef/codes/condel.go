package main

import (
	. "fmt"
	"io"
)

func conDel(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		total := 0
		sum := 0
		minS := int(1e9)
		for r, v := range a {
			total += v
			sum += v
			l := r + 1 - k
			if l < 0 {
				continue
			}
			minS = min(minS, sum)
			sum -= a[l]
		}
		Fprintln(out, minS*(minS+1)/2+total-minS)
	}
}

//func main() { conDel(bufio.NewReader(os.Stdin), os.Stdout) }
