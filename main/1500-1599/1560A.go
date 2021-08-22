package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1560A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for i := 1; ; i++ {
			if i%3 > 0 && i%10 != 3 {
				if n--; n == 0 {
					Fprintln(out, i)
					break
				}
			}
		}
	}
}

//func main() { CF1560A(os.Stdin, os.Stdout) }
