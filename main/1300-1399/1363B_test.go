package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1363B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
001
100
101
010
0
1
001100
outputCopy
0
0
1
1
0
0
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1363B)
}
