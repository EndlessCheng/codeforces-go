package _00_199

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF123D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aaaa
outputCopy
20
inputCopy
abcdef
outputCopy
21
inputCopy
abacabadabacaba
outputCopy
188`
	testutil.AssertEqualCase(t, rawText, 0, CF123D)
}
