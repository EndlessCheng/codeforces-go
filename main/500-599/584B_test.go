package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF584B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
20
inputCopy
2
outputCopy
680`
	testutil.AssertEqualCase(t, rawText, 0, CF584B)
}
