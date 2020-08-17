package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1332A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 2 2 2
0 0 -2 -2 2 2
3 1 4 1
0 0 -1 -1 1 1
1 1 1 1
1 1 1 1 1 1
0 0 0 1
0 0 0 0 0 1
5 1 1 1
0 0 -100 -100 0 100
1 1 5 1
0 0 -100 -100 100 0
outputCopy
Yes
No
No
Yes
Yes
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF1332A)
}
