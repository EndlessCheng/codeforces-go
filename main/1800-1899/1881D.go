package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1881D(in io.Reader, out io.Writer) {
	var T, n, x int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := map[int]int{}
		for range n {
			Fscan(in, &x)
			for i := 2; i*i <= x; i++ {
				if x%i > 0 {
					continue
				}
				e := 1
				for x /= i; x%i == 0; x /= i {
					e++
				}
				cnt[i] += e
			}
			if x > 1 {
				cnt[x]++
			}
		}
		for _, c := range cnt {
			if c%n > 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1881D(bufio.NewReader(os.Stdin), os.Stdout) }
