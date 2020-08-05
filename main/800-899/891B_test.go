package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF891B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
2 1 
inputCopy
4
1000 100 10 1
outputCopy
100 1 1000 10`
	testutil.AssertEqualCase(t, rawText, 0, CF891B)
}
