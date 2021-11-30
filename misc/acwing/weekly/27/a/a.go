package main

import (
	. "fmt"
	"io"
	"os"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for i, p10 := 1, 10; ; i++ {
			if i*p10 > n {
				Fprintln(out, strconv.Itoa(n / i)[n%i]&15)
				break
			}
			n += p10
			p10 *= 10
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
