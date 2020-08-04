package _00_399

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF349B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	v, min := 0, int(1e9)
	a := [10]int{}
	Fscan(in, &v)
	for i := 1; i <= 9; i++ {
		Fscan(in, &a[i])
		if a[i] < min {
			min = a[i]
		}
	}

	l := v / min
	if l == 0 {
		Fprint(out, -1)
		return
	}
	for ; l >= 0; l-- {
		for i := 9; i > 0; i-- {
			if c := a[i]; c <= v && (v-c)/min == l {
				v -= c
				Fprint(out, i)
				break
			}
		}
	}
}

//func main() { CF349B(os.Stdin, os.Stdout) }
