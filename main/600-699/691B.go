package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf691B(in io.Reader, out io.Writer) {
	r := [2]string{"$A------HI---M-O----TUVWXY-", "$-d-b----------oqp----vwx--"}
	s := ""
	Fscan(in, &s)
	for i, n := 0, len(s); i <= n/2; i++ {
		v := s[i]
		if r[v>>5&1][v&31] != s[n-1-i] {
			Fprint(out, "NIE")
			return
		}
	}
	Fprint(out, "TAK")
}

//func main() { cf691B(os.Stdin, os.Stdout) }
