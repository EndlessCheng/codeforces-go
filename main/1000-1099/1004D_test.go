package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/1004/D
// https://codeforces.com/problemset/status/1004/problem/D
func TestCF1004D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
20
1 0 2 3 5 3 2 1 3 2 3 1 4 2 1 4 2 3 2 4
outputCopy
4 5
2 2
inputCopy
18
2 2 3 2 4 3 3 3 0 2 4 2 1 3 2 1 1 1
outputCopy
3 6
2 3
inputCopy
6
2 1 0 2 1 2
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1004D)
}

func TestCheckCF1004D(t *testing.T) {
	//return
	assert := assert.New(t)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	inputGenerator := func() (string, testutil.OutputChecker) {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.IntOnly(1, 8)
		m := rg.IntOnly(n, 8)
		x, y := rg.IntOnly(0, (n-1)/2), rg.IntOnly(0, (m-1)/2)
		rg.Bytes(strconv.Itoa(n * m))
		rg.NewLine()
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				rg.Bytes(strconv.Itoa(abs(i-x)+abs(j-y)) + " ")
			}
		}
		return rg.String(), func(myOutput string) (_b bool) {
			// 检查 myOutput 是否符合题目要求
			// * 最好重新看一遍题目描述以免漏判 *
			// 对于 special judge 的题目，可能还需要额外跑个暴力来检查 myOutput 是否满足最优解等
			in := strings.NewReader(myOutput)

			var nn, mm, xx, yy int
			Fscan(in, &nn, &mm, &xx, &yy)
			if nn == -1 {
				panic(-1)
			}

			if !assert.EqualValues([4]int{n, m, x+1, y+1}, [4]int{nn, mm, xx, yy}) {
				return
			}

			return true
		}
	}

	testutil.CheckRunResultsInf(t, inputGenerator, CF1004D)
}
