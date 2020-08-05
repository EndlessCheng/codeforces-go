package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF550A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ABA
outputCopy
NO
inputCopy
BACFAB
outputCopy
YES
inputCopy
AXBYBXA
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF550A)
}
