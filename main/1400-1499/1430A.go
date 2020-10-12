package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1430A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for c7 := 0; c7*7 <= n; c7++ {
			for c5 := 0; c7*7+c5*5 <= n; c5++ {
				if (n-(c7*7+c5*5))%3 == 0 {
					Fprintln(out, (n-(c7*7+c5*5))/3, c5, c7)
					continue o
				}
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { CF1430A(os.Stdin, os.Stdout) }
