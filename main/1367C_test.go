package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1367C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6 1
100010
6 2
000000
5 1
10101
3 1
001
2 2
00
1 1
0
outputCopy
1
2
0
1
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1367C)
}
