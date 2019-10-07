package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol427D(t *testing.T) {
	// just copy from website
	rawText := `
apple
pepperoni
outputCopy
2
inputCopy
lover
driver
outputCopy
1
inputCopy
bidhan
roy
outputCopy
-1
inputCopy
testsetses
teeptes
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, Sol427D)
}
