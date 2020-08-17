package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1312D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
outputCopy
6
inputCopy
3 5
outputCopy
10
inputCopy
42 1337
outputCopy
806066790
inputCopy
100000 200000
outputCopy
707899035`
	testutil.AssertEqualCase(t, rawText, 0, CF1312D)
}
