package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1307D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5 3
1 3 5
1 2
2 3
3 4
3 5
2 4
outputCopy
3
inputCopy
5 4 2
2 4
1 2
2 3
3 4
4 5
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1307D)
}
