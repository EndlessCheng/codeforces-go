package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF559C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4 2
2 2
2 3
outputCopy
2
inputCopy
100 100 3
15 16
16 15
99 88
outputCopy
545732279`
	testutil.AssertEqualCase(t, rawText, 0, CF559C)
}
