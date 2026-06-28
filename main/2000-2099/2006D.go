package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2006D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q, &k)
		sq := 0
		for i := 1; i*i <= k; i++ {
			sq = i
		}
		psum := make([]int, n+1)
		csum := make([]int, q)
		ccnt := make([]int, q)
		l := make([]int, q)
		r := make([]int, q)
		ans := make([]int, q)
		v1 := make([][]int, sq+1)
		v2 := make([][]int, sq+1)
		for i := range n {
			var x int
			Fscan(in, &x)
			if x <= sq {
				v1[x] = append(v1[x], i)
			} else {
				v2[k/x] = append(v2[k/x], i)
			}
		}

		for i := range q {
			Fscan(in, &l[i], &r[i])
			l[i]--
		}

		for i := 1; i <= sq; i++ {
			clear(psum)
			for _, p := range v1[i] {
				psum[p+1]--
			}
			for _, p := range v2[i] {
				psum[p+1]++
			}
			for j := 1; j <= n; j++ {
				psum[j] += psum[j-1]
			}
			for j := 0; j < q; j++ {
				csum[j] += psum[r[j]]
				csum[j] -= psum[l[j]]
			}

			clear(psum)
			for _, p := range v1[i] {
				psum[p+1]++
			}
			for _, p := range v2[i] {
				psum[p+1]++
			}
			for j := 1; j <= n; j++ {
				psum[j] += psum[j-1]
			}
			for j := range q {
				ccnt[j] += psum[r[j]]
				ccnt[j] -= psum[l[j]]
				if ccnt[j] == r[j]-l[j] {
					ans[j] = max(ans[j], csum[j]/2)
				} else {
					ans[j] = max(ans[j], (csum[j]+1)/2)
				}
			}
		}

		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2006D(bufio.NewReader(os.Stdin), os.Stdout) }
