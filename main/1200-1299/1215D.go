package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1215D(in io.Reader, out io.Writer) {
	var n, sl, cl, sr, cr int
	var s []byte
	Fscan(bufio.NewReader(in), &n, &s)
	for _, b := range s[:n/2] {
		if b != '?' {
			sl += int(b & 15)
		} else {
			cl++
		}
	}
	for _, b := range s[n/2:] {
		if b != '?' {
			sr += int(b & 15)
		} else {
			cr++
		}
	}
	if (cl+cr)&1 > 0 {
		Fprint(out, "Monocarp")
		return
	}
	if sl == sr {
		if cl != cr {
			Fprint(out, "Monocarp")
		} else {
			Fprint(out, "Bicarp")
		}
		return
	}
	if cl > cr {
		sl, cl, sr, cr = sr, cr, sl, cl
	}
	if sl != sr+(cr-cl)/2*9 {
		Fprint(out, "Monocarp")
	} else {
		Fprint(out, "Bicarp")
	}
}

//func main() { CF1215D(os.Stdin, os.Stdout) }
