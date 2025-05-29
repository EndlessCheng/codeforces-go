package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf508B(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	n := len(s)
	last := -1
	for i, b := range s {
		if b%2 > 0 {
			continue
		}
		last = i
		if b < s[n-1] {
			break
		}
	}
	if last < 0 {
		Fprint(out, -1)
		return
	}
	s[last], s[n-1] = s[n-1], s[last]
	Fprintf(out, "%s", s)
}

//func main() { cf508B(bufio.NewReader(os.Stdin), os.Stdout) }
