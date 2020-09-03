package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF533C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 1 2 2
outputCopy
Polycarp
inputCopy
4 7 7 4
outputCopy
Vasiliy`
	testutil.AssertEqualCase(t, rawText, 0, CF533C)
}
