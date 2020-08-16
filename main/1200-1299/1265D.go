package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1265D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	n := 0
	aa := [2][4]int{}
	for i := range aa[0] {
		Fscan(in, &aa[0][i])
		aa[1][3-i] = aa[0][i]
		n += aa[0][i]
	}
	for order, a := range aa {
		even := a[0] + a[2]
		odd := n - even
		if even != odd && even-1 != odd {
			continue
		}
		b := make([]int, n)
		for i := 0; i < n; i += 2 {
			if a[0] > 0 {
				a[0]--
			} else if a[2] > 0 {
				a[2]--
				b[i] = 2
			}
		}
		for i := 1; i < n; i += 2 {
			if a[1] > 0 {
				a[1]--
				b[i] = 1
			} else if a[3] > 0 {
				a[3]--
				b[i] = 3
			}
		}
		ok := true
		for i := 1; i < n; i++ {
			if b[i]-b[i-1] != 1 && b[i]-b[i-1] != -1 {
				ok = false
				break
			}
		}
		if ok {
			Fprintln(out, "YES")
			for _, v := range b {
				if order == 0 {
					Fprint(out, v, " ")
				} else {
					Fprint(out, 3-v, " ")
				}
			}
			return
		}
	}
	Fprint(out, "NO")
}

//func main() {
//	Sol1265D(os.Stdin, os.Stdout)
//}
