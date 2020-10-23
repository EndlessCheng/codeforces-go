package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF518A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
a
c
outputCopy
b
inputCopy
aaa
zzz
outputCopy
kkk
inputCopy
abcdefg
abcdefh
outputCopy
No such string`
	testutil.AssertEqualCase(t, rawText, 0, CF518A)
}
