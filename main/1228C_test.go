package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1228C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 2
outputCopy
2
inputCopy
20190929 1605
outputCopy
363165664
inputCopy
947 987654321987654321
outputCopy
593574252`
	testutil.AssertEqualCase(t, rawText, 0, CF1228C)
}
