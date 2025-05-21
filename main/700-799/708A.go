package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf708A(in io.Reader, out io.Writer) {
	var t []byte
	Fscan(in, &t)
	for i, c := range t {
		if c > 'a' {
			for ; i < len(t) && t[i] > 'a'; i++ {
				t[i]--
			}
			Fprintf(out, "%s", t)
			return
		}
	}
	t[len(t)-1] = 'z'
	Fprintf(out, "%s", t)
}

//func main() { cf708A(bufio.NewReader(os.Stdin), os.Stdout) }
