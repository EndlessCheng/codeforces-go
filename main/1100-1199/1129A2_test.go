package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

// http://codeforces.com/problemset/problem/1129/A2
// https://codeforces.com/problemset/status/1129/problem/A2
func TestCF1129A2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 7
2 4
5 1
2 3
3 4
4 1
5 3
3 5
outputCopy
10 9 10 10 9 
inputCopy
2 3
1 2
1 2
1 2
outputCopy
5 6 
inputCopy
50 20
4 18
39 33
49 32
7 32
38 1
46 11
8 1
3 31
30 47
24 16
33 5
5 21
3 48
13 23
49 50
18 47
40 32
9 23
19 39
25 12
outputCopy
99 98 97 127 126 125 124 123 122 121 120 119 118 117 116 115 114 113 112 111 110 109 108 107 106 105 104 103 102 101 100 99 98 97 96 95 94 93 92 93 92 91 90 89 88 87 86 85 84 100
inputCopy
4 3
1 4
2 3
2 3
outputCopy
6 6 8 7`
	testutil.AssertEqualCase(t, rawText, 0, CF1129A2)
}

func TestCompareCF1129A2(t *testing.T) {
	return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 10)
		m := rg.Int(1, 10)
		rg.NewLine()
		for i := 0; i < m; i++ {
			a := rg.Int(1, n-1)
			b := rg.Int(a+1, n)
			if rand.Float64() < 0.5 {
				a, b = b, a
			}
			rg.NewLine()
		}
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF1129A2, CF1129A2)
}

func distCF1129A2(a, b, n int) int {
	return (b + n - a) % n
}

func runBF1129A2(in io.Reader, out io.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	nums, cnt := [5004]int{}, [5004]int{}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		if nums[a] == 0 {
			nums[a] = b
		} else {
			cnt[a] += 1
			if distCF1129A2(a, b, n) < distCF1129A2(a, nums[a], n) {
				nums[a] = b
			}

		}
	}
	res := ""
	for i := 1; i < n+1; i++ {
		ans := 0
		for j := 1; j < n+1; j++ {
			if nums[j] > 0 {
				ans = int(math.Max(float64(ans), float64(distCF1129A2(i, j, n)+cnt[j]*n+distCF1129A2(j, nums[j], n))))
			}
		}
		res += strconv.Itoa(ans) + " "
	}
	fmt.Fprintln(out, res[:len(res)-1])
}
