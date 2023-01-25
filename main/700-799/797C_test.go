package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/797/C
// https://codeforces.com/problemset/status/797/problem/C
func TestCF797C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
cab
outputCopy
abc
inputCopy
acdb
outputCopy
abdc
inputCopy
eeddcbeeec
outputCopy
bcceeeddee
inputCopy
cad
outputCopy
acd`
	testutil.AssertEqualCase(t, rawText, 0, CF797C)
}

func TestCompareCF797C(t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.Str(1, 5, 'a', 'd')
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		min := func(a, b byte) byte {
			if a > b {
				return b
			}
			return a
		}
		totalLen, counterResult, counterTotal, counterCompare := 0,0,0,0
		inputString := ""
		compareString := [100001]byte{}
		resultString := [100001]byte{}
		e := []byte{124}
		Fscan(in, &inputString)
		totalLen = len(inputString)
		compareString[totalLen] = e[0]

		for i := totalLen -1;i>-1; i-- {
			compareString[i] = min(compareString[i + 1], inputString[i])
		}

		for counterTotal < totalLen {
			if len(inputString)< counterCompare || len(resultString)< counterResult {
				break
			}
			resultString[counterResult] = inputString[counterCompare]
			counterResult++
			counterCompare++
			for counterResult >0 && resultString[counterResult-1] <= compareString[counterCompare] {
				counterResult--
				Fprint(out, string(resultString[counterResult]))
				counterTotal++
			}
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF797C)
}
