package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1303D(t *testing.T) {
	// just copy from website
	rawText := `
3
10 3
1 32 1
23 4
16 1 4 1
20 5
2 1 16 1 8
outputCopy
2
-1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1303D)
}
