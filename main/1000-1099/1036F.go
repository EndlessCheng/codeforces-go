package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1036F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	isSQ := func(x int64) bool {
		res := int64(math.Sqrt(float64(x)))
		return res*res == x
	}
	floorSqrt := func(x int64) int64 {
		res := int64(math.Sqrt(float64(x)))
		if res*res > x {
			res--
		}
		return res
	}

	// 生成 [2,1e18] 范围内的立方及以上的数
	a := make([]int64, 0, 1003479)
	for i := int64(2); i <= 1e6; i++ {
		for x := i * i; x <= 1e18/i; {
			if x *= i; !isSQ(x) {
				a = append(a, x)
			}
		}
	}
	// 排序+去重
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	k := 0
	for _, w := range a[1:] {
		if a[k] != w {
			k++
			a[k] = w
		}
	}
	a = a[:k+1]

	var T, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, n-int64(sort.Search(len(a), func(i int) bool { return a[i] > n }))-floorSqrt(n))
	}
}

//func main() { CF1036F(os.Stdin, os.Stdout) }
