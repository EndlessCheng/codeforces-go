package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1956E2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		for i := range n {
			Fscan(in, &a[i])
		}
		for range 2000 {
			for i := range n - 1 {
				a[i+1] = max(a[i+1]-a[i], 0)
			}
			a[0] = max(a[0]-a[n-1], 0)
		}

		d := 0
		for a[0] > 0 {
			a[1] = max(a[1]-a[0], 0)
			tmp := a[0]
			copy(a, a[1:n])
			a[n-1] = tmp
			d++
		}

		f := make([]int, n)
		cnt := 0
		for i := 1; i < n; {
			if a[i] > 0 {
				j := i
				for a[j] > 0 {
					j++
				}
				f[(i+d)%n] = 1
				cnt++
				if j-i == 3 && a[i+2] > ((a[i+1]%a[i]+a[i+1]-a[i])*(a[i+1]/a[i]))>>1 {
					f[(i+2+d)%n] = 1
					cnt++
				}
				i = j
			} else {
				i++
			}
		}

		Fprintln(out, cnt)
		for i, v := range f {
			if v > 0 {
				Fprint(out, i+1, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1956E2(bufio.NewReader(os.Stdin), os.Stdout) }
