package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1365B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
10 20 20 30
0 1 0 1
3
3 1 2
0 1 1
4
2 2 4 8
1 1 1 1
3
5 15 4
0 0 0
4
20 10 100 50
1 0 0 1
outputCopy
Yes
Yes
Yes
No
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF1365B)
}
