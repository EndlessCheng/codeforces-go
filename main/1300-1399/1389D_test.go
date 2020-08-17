package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1389D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 5
1 2
3 4
2 1000000000
1 1
999999999 999999999
10 3
5 10
7 8
outputCopy
7
2000000000
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1389D)
}
