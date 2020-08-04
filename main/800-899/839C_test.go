package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF839C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
1 3
2 4
outputCopy
1.500000000000000
inputCopy
5
1 2
1 3
3 4
2 5
outputCopy
2.000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF839C)
}
