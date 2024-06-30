package main

import (
	. "fmt"
	"io"
)

func cf1955E(in io.Reader, out io.Writer) {
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		diff := make([]byte, n+1)
	o:
		for k := n; k > 0; k-- {
			clear(diff)
			sd := byte('0')
			for i, b := range s {
				sd ^= diff[i]
				if sd != b {
					continue
				}
				if i+k > n {
					continue o
				}
				sd ^= 1
				diff[i+k] = 1
			}
			Fprintln(out, k)
			break
		}
	}
}

//func main() { cf1955E(bufio.NewReader(os.Stdin), os.Stdout) }
