package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1440A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 1 1 1
100
5 10 100 1
01010
5 10 1 1
11111
5 1 10 1
11111
12 2 1 10
101110110101
2 100 1 10
00
outputCopy
3
52
5
10
16
22`
	testutil.AssertEqualCase(t, rawText, 0, CF1440A)
}
