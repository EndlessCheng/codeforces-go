package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1076C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, D int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &D)
		if 0 < D && D < 4 {
			Fprintln(out, "N")
			continue
		}
		d := float64(D)
		dt := math.Sqrt(d * (d - 4))
		Fprintf(out, "Y %.9f %.9f\n", (d+dt)/2, (d-dt)/2)
	}
}

//func main() { CF1076C(os.Stdin, os.Stdout) }
