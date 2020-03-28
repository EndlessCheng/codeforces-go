package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1183A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
432
outputCopy
435
inputCopy
99
outputCopy
103
inputCopy
237
outputCopy
237
inputCopy
42
outputCopy
44`
	testutil.AssertEqualCase(t, rawText, 0, CF1183A)
}
