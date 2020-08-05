package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF915C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
123
222
outputCopy
213
inputCopy
3921
10000
outputCopy
9321
inputCopy
4940
5000
outputCopy
4940`
	testutil.AssertEqualCase(t, rawText, 0, CF915C)
}
