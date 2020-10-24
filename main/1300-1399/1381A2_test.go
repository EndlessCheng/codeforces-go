package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1381A2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2
01
10
5
01011
11100
2
01
01
10
0110011011
1000110100
1
0
1
outputCopy
3 1 2 1
6 5 2 5 3 1 2
0
9 4 1 2 10 4 1 2 1 5
1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1381A2)
}
