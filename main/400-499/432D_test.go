package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF432D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ABACABA
outputCopy
3
1 4
3 2
7 1
inputCopy
AAA
outputCopy
3
1 3
2 2
3 1`
	testutil.AssertEqualCase(t, rawText, 0, CF432D)
}
