package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF877A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
Alex_and_broken_contest
outputCopy
NO
inputCopy
NikitaAndString
outputCopy
YES
inputCopy
Danil_and_Olya
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF877A)
}
