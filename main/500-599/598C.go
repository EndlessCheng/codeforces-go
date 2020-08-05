package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF598C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct {
		a float64
		i int
	}

	var n, x, y, minI int
	Fscan(in, &n)
	a := make([]pair, n, n+1)
	for i := range a {
		Fscan(in, &x, &y)
		a[i] = pair{math.Atan2(float64(y), float64(x)), i + 1}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].a < a[j].a }) // FIXME: WA114
	a = append(a, pair{a[0].a + 2*math.Pi, a[0].i})
	for i := 1; i < n; i++ {
		if a[i+1].a+a[minI].a < a[minI+1].a+a[i].a {
			minI = i
		}
	}
	Fprint(out, a[minI].i, a[minI+1].i)
}

//func main() { CF598C(os.Stdin, os.Stdout) }
