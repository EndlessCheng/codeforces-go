package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1323B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	makeSum := func(n int) ([]int, []int) {
		sum1 := 0
		times := make([]int, n+1)
		for i := 0; i < n; i++ {
			var v int
			if Fscan(in, &v); v == 1 {
				sum1++
			} else {
				if sum1 > 0 {
					times[sum1]++
				}
				sum1 = 0
			}
		}
		if sum1 > 0 {
			times[sum1]++
		}
		sumLen := make([]int, n+1)
		sumTimes := make([]int, n+1)
		for i := 1; i <= n; i++ {
			sumLen[i] = sumLen[i-1] + times[i]*i
			sumTimes[i] = sumTimes[i-1] + times[i]
		}
		return sumLen, sumTimes
	}

	var n, m, k int
	Fscan(in, &n, &m, &k)
	sumLenA, sumTimesA := makeSum(n)
	sumLenB, sumTimesB := makeSum(m)
	ans := int64(0)
	for x := n; x > 0; x-- {
		if k%x == 0 {
			y := k / x
			if y > m {
				break
			}
			cx := int64(sumLenA[n] - sumLenA[x-1] - (sumTimesA[n]-sumTimesA[x-1])*(x-1))
			cy := int64(sumLenB[m] - sumLenB[y-1] - (sumTimesB[m]-sumTimesB[y-1])*(y-1))
			ans += cx * cy
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1323B(os.Stdin, os.Stdout) }
