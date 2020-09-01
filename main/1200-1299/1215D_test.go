package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1215D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0523
outputCopy
Bicarp
inputCopy
2
??
outputCopy
Bicarp
inputCopy
8
?054??0?
outputCopy
Bicarp
inputCopy
6
???00?
outputCopy
Monocarp`
	testutil.AssertEqualCase(t, rawText, 0, CF1215D)
}
