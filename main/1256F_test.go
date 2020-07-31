package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1256F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
abcd
abdc
5
ababa
baaba
4
asdf
asdg
4
abcd
badc
outputCopy
NO
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1256F)
}
