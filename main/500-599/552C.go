package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol552C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var w, m int
	Fscan(in, &w, &m)
	if w <= 3 {
		Fprint(out, "YES")
		return
	}
	coef := []int{}
	for ; m > 0; m /= w {
		c := m % w
		if c == w-1 {
			c = -1
		}
		coef = append(coef, c)
	}
	n := len(coef)
	for i := 0; i < n; i++ {
		c := coef[i]
		if c == -1 {
			for j := i + 1; j < n; j++ {
				coef[j]++
				if coef[j] != 0 {
					break
				}
			}
		} else if c > 1 {
			if c < w-1 {
				Fprint(out, "NO")
				return
			}
			coef[i] = -1
			i--
		}
	}
	Fprint(out, "YES")
}

//func main() {
//	Sol552C(os.Stdin, os.Stdout)
//}