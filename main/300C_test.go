package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF300C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 3 3
outputCopy
1
inputCopy
2 3 10
outputCopy
165`
	testutil.AssertEqualCase(t, rawText, 0, CF300C)
}
