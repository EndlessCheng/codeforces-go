package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF466B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 5
outputCopy
18
3 6
inputCopy
2 4 4
outputCopy
16
4 4`
	testutil.AssertEqualCase(t, rawText, 0, CF466B)
}
