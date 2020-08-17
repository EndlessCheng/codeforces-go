package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1368D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
123
outputCopy
15129
inputCopy
3
1 3 5
outputCopy
51
inputCopy
2
349525 699050
outputCopy
1099509530625`
	testutil.AssertEqualCase(t, rawText, 0, CF1368D)
}
