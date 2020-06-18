package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF16C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
800 600 4 3
outputCopy
800 600
inputCopy
1920 1200 16 9
outputCopy
1920 1080
inputCopy
1 1 1 2
outputCopy
0 0`
	testutil.AssertEqualCase(t, rawText, 0, CF16C)
}
