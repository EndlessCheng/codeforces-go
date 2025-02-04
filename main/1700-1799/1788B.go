package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1788B(in io.Reader, out io.Writer) {
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		a := [2][]byte{}
		j := 0
		for i := range s {
			c := s[i] - '0'
			a[0] = append(a[0], '0'+c/2)
			a[1] = append(a[1], '0'+c/2)
			if c%2 > 0 {
				a[j][len(a[j])-1]++
				j ^= 1
			}
		}
		for len(a[1]) > 1 && a[1][0] == '0' {
			a[1] = a[1][1:]
		}
		Fprintf(out, "%s %s\n", a[0], a[1])
	}
}

//func main() { cf1788B(bufio.NewReader(os.Stdin), os.Stdout) }
