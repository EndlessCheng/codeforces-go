package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1082D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 2 2
outputCopy
YES 2
2
1 2
2 3
inputCopy
5
1 4 1 1 1
outputCopy
YES 2
4
1 2
3 2
4 2
5 2
inputCopy
3
1 1 1
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1082D)
}
