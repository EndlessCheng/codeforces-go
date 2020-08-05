package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func Sol935C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var flatR, x1, y1, x2, y2 float64
	Fscan(in, &flatR, &x1, &y1, &x2, &y2)
	x2 -= x1
	y2 -= y1
	if x2 == 0 && y2 == 0 {
		Fprint(out, x1+flatR/2, y1, flatR/2)
		return
	}
	if x2*x2+y2*y2 >= flatR*flatR {
		Fprint(out, x1, y1, flatR)
		return
	}
	r2 := math.Hypot(x2, y2)
	r := (r2 + flatR) / 2
	k := (r - r2) / r2
	Fprint(out, x1-k*x2, y1-k*y2, r)
}

//func main() {
//	Sol935C(os.Stdin, os.Stdout)
//}
