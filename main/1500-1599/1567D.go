package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func CF1567D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, s, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &n)
		x := int(math.Pow10(len(strconv.Itoa(s)) - 1))
		for i := 1; i < n; i++ {
			for s-x < n-i {
				x /= 10
			}
			Fprint(out, x, " ")
			s -= x
		}
		Fprintln(out, s)
	}
}

//func main() { CF1567D(os.Stdin, os.Stdout) }
