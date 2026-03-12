package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf2143E(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		sum := 0
		for i, b := range s {
			if b == '(' {
				sum += 1 - i%2*2
			}
		}
		if n%2 != 0 || (sum-n/2)%2 != 0 || sum == -n/2 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, strings.Repeat("()", (sum+n/2)/2-1)+"("+strings.Repeat("()", (n/2-sum)/2)+")")
		}
	}
}

//func main() { cf2143E(bufio.NewReader(os.Stdin), os.Stdout) }
