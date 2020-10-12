package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1430D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
6
111010
1
0
1
1
2
11
6
101010
outputCopy
3
1
1
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1430D)
}
