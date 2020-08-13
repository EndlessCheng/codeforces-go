package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF615D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 3
outputCopy
36
inputCopy
3
2 3 2
outputCopy
1728`
	testutil.AssertEqualCase(t, rawText, 0, CF615D)
}