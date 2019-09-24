package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol584D(t *testing.T) {
	// just copy from website
	rawText := `
27
outputCopy
3
5 11 11`
	testutil.AssertEqual(t, rawText, Sol584D)
}
