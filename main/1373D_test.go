package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1373D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
8
1 7 3 4 7 6 2 9
5
1 2 1 2 1
10
7 8 4 5 7 6 8 9 7 3
4
3 1 2 1
outputCopy
26
5
37
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1373D)
}
