package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1567B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		xor := 0
		switch a % 4 {
		case 1:
			xor = a - 1
		case 2:
			xor = 1
		case 3:
			xor = a
		}
		if xor == b {
			Fprintln(out, a)
		} else if xor^b != a {
			Fprintln(out, a+1)
		} else {
			Fprintln(out, a+2)
		}
	}
}

//func main() { CF1567B(os.Stdin, os.Stdout) }
