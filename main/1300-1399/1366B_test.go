package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1366B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6 4 3
1 6
2 3
5 5
4 1 2
2 4
1 2
3 3 2
2 3
1 2
outputCopy
6
2
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1366B)
}
