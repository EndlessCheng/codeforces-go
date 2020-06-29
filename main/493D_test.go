package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF493D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
white
1 2
inputCopy
3
outputCopy
black`
	testutil.AssertEqualCase(t, rawText, 0, CF493D)
}
