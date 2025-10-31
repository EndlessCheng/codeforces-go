package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func p12751(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, num, k int
	Fscan(in, &n, &m)
	a := make([]int, n)
	B := int(math.Sqrt(float64(m)))
	type pair struct{ l, num int }
	groups := make([][]pair, B)
	for ; m > 0; m-- {
		Fscan(in, &l, &num, &k)
		l--
		if k < B {
			groups[k] = append(groups[k], pair{l, num})
		} else {
			for i := 0; i < num; i++ {
				a[l+i*k]++
			}
		}
	}

	diff := make([]int, n+1)
	for k, g := range groups {
		if g == nil {
			continue
		}
		buckets := make([][]pair, k)
		for _, p := range g {
			buckets[p.l%k] = append(buckets[p.l%k], p)
		}
		for start, bucket := range buckets {
			if bucket == nil {
				continue
			}
			if len(bucket) == 1 {
				p := bucket[0]
				for i := 0; i < p.num; i++ {
					a[p.l+i*k]++
				}
				continue
			}

			m := (n - start - 1) / k
			for i := 0; i <= m; i++ {
				diff[i] = 0
			}
			for _, p := range bucket {
				diff[p.l/k]++
				diff[p.l/k+p.num]--
			}

			sd := 0
			for i := 0; i <= m; i++ {
				sd += diff[i]
				a[start+i*k] += sd
			}
		}
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
	Fprintln(out)
}

//func main() { p12751(bufio.NewReader(os.Stdin), os.Stdout) }
