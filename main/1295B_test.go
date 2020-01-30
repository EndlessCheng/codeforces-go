package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1295B(t *testing.T) {
	// just copy from website
	rawText := `
4
6 10
010010
5 3
10101
1 0
0
2 0
01
outputCopy
3
0
1
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1295B)
}
