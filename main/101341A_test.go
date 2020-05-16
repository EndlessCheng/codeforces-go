package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF101341A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
(
()
)
outputCopy
YES
1 2 3
inputCopy
3
)))
((
(
outputCopy
YES
2 3 1
inputCopy
3
)))
((
)
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF101341A)
}
