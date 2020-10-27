package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF877A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	c := 0
	for _, name := range []string{"Danil", "Olya", "Slava", "Ann", "Nikita"} {
		c += strings.Count(s, name)
	}
	if c == 1 {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}
}

//func main() { CF877A(os.Stdin, os.Stdout) }
