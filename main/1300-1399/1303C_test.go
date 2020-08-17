package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1303C(t *testing.T) {
	// just copy from website
	rawText := `
5
ababa
codedoca
abcda
zxzytyz
abcdefghijklmnopqrstuvwxyza
outputCopy
YES
bacdefghijklmnopqrstuvwxyz
YES
edocabfghijklmnpqrstuvwxyz
NO
YES
xzytabcdefghijklmnopqrsuvw
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1303C)
}
