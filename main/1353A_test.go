package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1353A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 100
2 2
5 5
2 1000000000
1000000000 1000000000
outputCopy
0
2
10
1000000000
2000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1353A)
}
