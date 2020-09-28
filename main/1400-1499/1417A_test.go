package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1417A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 2
1 1
3 5
1 2 3
3 7
3 2 2
outputCopy
1
5
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1417A)
}
