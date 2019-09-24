package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
	)

func TestSol1202D(t *testing.T) {
	// just copy from website
	rawText := `2
6
1
outputCopy
113337
1337`
	testutil.AssertEqual(t, rawText, Sol1202D)
}
