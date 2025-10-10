package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2121E(in io.Reader, out io.Writer) {
	var T int
	var l, r []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		for i, d := range l {
			if d == r[i] {
				continue
			}
			if r[i]-d > 1 {
				Fprintln(out, i*2)
			} else {
				j := i + 1
				for ; j < len(l) && l[j] == '9' && r[j] == '0'; j++ {
				}
				Fprintln(out, i+j)
			}
			continue o
		}
		Fprintln(out, len(l)*2)
	}
}

//func main() { cf2121E(bufio.NewReader(os.Stdin), os.Stdout) }
