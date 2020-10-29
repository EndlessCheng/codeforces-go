package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1369D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1
2
3
4
5
100
2000000
outputCopy
0
0
4
4
12
990998587
804665184`
	testutil.AssertEqualCase(t, rawText, 0, CF1369D)
}
