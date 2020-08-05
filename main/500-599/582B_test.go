package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"sort"
	"testing"
)

func TestCF582B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
3 1 4 2
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, -1, CF582B)
}

// 无尽对拍
func TestCF582BCmp(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		rg := testutil.NewRandGenerator()
		n := rg.Int(100, 100)
		rg.Int(200, 400)
		rg.NewLine()
		rg.IntSlice(n, 1, 300)
		//Println(rg.String())
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var n, t int
		Fscan(in, &n, &t)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := []int{}
		for i := 0; i < t; i++ {
			b = append(b, a...)
		}
		a = b
		n = len(a)
		dp := make([]int, 0, n)
		for _, v := range a {
			if i := sort.SearchInts(dp, v+1); i < len(dp) {
				dp[i] = v
			} else {
				dp = append(dp, v)
			}
		}
		Fprint(out, len(dp))
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF582B)
}
