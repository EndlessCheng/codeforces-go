package main

import (
	. "fmt"
	"io"
)

func cf1895C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := [6][]string{}
	for i := 0; i < n; i++ {
		var s string
		Fscan(in, &s)
		a[len(s)] = append(a[len(s)], s)
	}
	
	ans := 0
	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {
			if (i+j)%2 > 0 {
				continue
			}
			m := (i + j) / 2
			cnt := [100]int{}
			for _, s := range a[i] {
				sum := 50
				for k, b := range s {
					if k < m {
						sum += int(b - '0')
					} else {
						sum -= int(b - '0')
					}
				}
				cnt[sum]++
			}
			for _, s := range a[j] {
				sum := 50
				for k, b := range s {
					if i+k < m {
						sum -= int(b - '0')
					} else {
						sum += int(b - '0')
					}
				}
				ans += cnt[sum]
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1895C(bufio.NewReader(os.Stdin), os.Stdout) }
