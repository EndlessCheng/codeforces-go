package main

import (
	"bytes"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/bits"
	"testing"
)

// https://codeforces.com/problemset/problem/476/E
// https://codeforces.com/problemset/status/476/problem/E
func TestCF476E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aaaaa
aa
outputCopy
2 2 1 1 0 0
inputCopy
axbaxxb
ab
outputCopy
0 1 1 2 1 1 0 0
inputCopy
bc
cc
outputCopy
0 0 0`
	testutil.AssertEqualCase(t, rawText, -1, CF476E)
}

func TestCompareCF476E(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.Str(1, 7, 'a', 'e')
		rg.NewLine()
		rg.Str(1, 7, 'a', 'e')
		return rg.String()
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var s, p []byte
		Fscan(in, &s, &p)

		f := func(s []byte) (c int) {
			for {
				i := bytes.Index(s, p)
				if i < 0 {
					break
				}
				c++
				s = s[i+len(p):]
			}
			return
		}

		ans := make([]int, len(s)+1)
		calc := func(sub uint)  {
			del := len(s) - bits.OnesCount(sub)
			ss := []byte{}
			for ; sub > 0; sub &= sub - 1 {
				p := bits.TrailingZeros(sub)
				v := s[p]
				ss = append(ss, v)
			}
			ans[del] = max(ans[del], f(ss))
		}
		for sub := uint(0); sub < 1<<len(s); sub++ {
			 calc(sub)
		}

		for _, v := range ans {
			Fprint(out, v, " ")
		}
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
//	rawText := `
//inputCopy
//aaaaa
//aa
//outputCopy
//2 2 1 1 0 0
//inputCopy
//axbaxxb
//ab
//outputCopy
//0 1 1 2 1 1 0 0`
//	testutil.AssertEqualCase(t, rawText, 0, runBF)
//	return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF476E)
}
