package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2168A2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var tp, s string
	Fscan(in, &tp)
	if tp[0] == 'f' {
		var n, v int
		Fscan(in, &n)
		ans := make([]byte, 0, n*8)
		for range n {
			Fscan(in, &v)
			for i := 28; i >= 0; i -= 4 {
				ans = append(ans, 'a'+byte(v>>i&15))
			}
		}
		Fprintf(out, "%s", ans)
	} else {
		Fscan(in, &s)
		Fprintln(out, len(s)/8)
		for i := 8; i <= len(s); i += 8 {
			v := 0
			for _, b := range s[i-8 : i] {
				v = v<<4 | int(b-'a')
			}
			Fprint(out, v, " ")
		}
	}
}

//func main() { cf2168A2(bufio.NewReader(os.Stdin), os.Stdout) }
