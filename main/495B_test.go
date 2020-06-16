package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF495B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
21 5
outputCopy
2
inputCopy
9435152 272
outputCopy
282
inputCopy
10 10
outputCopy
infinity`
	testutil.AssertEqualCase(t, rawText, 0, CF495B)
}
