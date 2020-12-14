package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF353D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var s []byte
	Fscan(in, &s)
	ans, m := 0, 0
	for _, b := range s {
		if b == 'M' {
			m++
		} else if m > 0 {
			ans = max(ans+1, m)
		}
	}
	Fprint(out, ans)
}

//func main() { CF353D(os.Stdin, os.Stdout) }
