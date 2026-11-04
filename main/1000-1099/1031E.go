package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1031E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	ans := [][2]int{}

	for i := 1; i <= n-2; i++ {
		if a[i] != 0 {
			k := 1
			for j := 2; j <= (n-i)/2; j++ {
				if a[i+j] != 0 {
					k = j
					break
				}
			}
			a[i] ^= 1
			a[i+k] ^= 1
			a[i+k+k] ^= 1
			ans = append(ans, [2]int{i, k})
		}
	}

	for i := n - 1; i <= n; i++ {
		if a[i] != 0 {
			if i <= 6 {
				Fprint(out, "NO")
				return
			}
			ans = append(ans, [2]int{i - 6, 1})
			ans = append(ans, [2]int{i - 6, 3})
			ans = append(ans, [2]int{i - 5, 1})
		}
	}

	Fprintln(out, "YES")
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprintln(out, v[0], v[0]+v[1], v[0]+v[1]+v[1])
	}
}

//func main() { cf1031E(bufio.NewReader(os.Stdin), os.Stdout) }
