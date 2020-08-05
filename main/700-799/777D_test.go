package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol777D(t *testing.T) {
	// just copy from website
	rawText := `
3
#book
#bigtown
#big
outputCopy
#b
#big
#big
inputCopy
3
#book
#cool
#cold
outputCopy
#book
#co
#cold
inputCopy
4
#car
#cart
#art
#at
outputCopy
#
#
#art
#at
inputCopy
3
#apple
#apple
#fruit
outputCopy
#apple
#apple
#fruit`
	testutil.AssertEqualCase(t, rawText, -1, Sol777D)
}
