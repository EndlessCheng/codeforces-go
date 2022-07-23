package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/725/C
// https://codeforces.com/problemset/status/725/problem/C
func TestCF725C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ABCDEFGHIJKLMNOPQRSGTUVWXYZ
outputCopy
YXWVUTGHIJKLM
ZABCDEFSRQPON
inputCopy
BUVTYZFQSNRIWOXXGJLKACPEMDH
outputCopy
Impossible
inputCopy
UTEDBZRVWLOFUASHCYIPXGJMKNQ
outputCopy
PIYCHSAUTEDBZ
XGJMKNQFOLWVR`
	testutil.AssertEqualCase(t, rawText, 0, CF725C)
}
