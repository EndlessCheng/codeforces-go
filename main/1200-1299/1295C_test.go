package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1295C(t *testing.T) {
	// just copy from website
	rawText := `
3
aabce
ace
abacaba
aax
ty
yyt
outputCopy
1
-1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1295C)
}
