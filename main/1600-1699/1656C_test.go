package main

import (
	"bufio"
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"sort"
	"testing"
)

// https://codeforces.com/problemset/problem/1656/C
// https://codeforces.com/problemset/status/1656/problem/C
func TestCF1656C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
2 5 6 8
3
1 1 1
5
4 1 7 0 8
4
5 9 17 5
outputCopy
YES
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1656C)
}

func TestCompareCF1656C(t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 99)
		rg.NewLine()
		rg.IntSlice(n, 0, 5)
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, CF1656C_AC, CF1656C)
}

func CF1656C_AC(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()


	var T int
	fmt.Fscan(in, &T)
	for t := 0; t < T; t++ {
		var n int
		fmt.Fscan(in, &n)

		var hasOne bool
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])

			if arr[i] == 1 {
				hasOne = true
			}
		}

		sort.Ints(arr)
		var f bool
		for i := 1; i < n; i++ {
			if arr[i]-arr[i-1] == 1 {
				f = true
				break
			}
		}

		if !hasOne || (hasOne && !f) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}