package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF630F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
outputCopy
29`
	testutil.AssertEqualCase(t, rawText, 0, CF630F)
}
