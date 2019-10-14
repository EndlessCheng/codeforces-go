package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol490C(t *testing.T) {
	// just copy from website
	rawText := `
116401024
97 1024
outputCopy
YES
11640
1024
inputCopy
284254589153928171911281811000
1009 1000
outputCopy
YES
2842545891539
28171911281811000
inputCopy
120
12 1
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, -1, Sol490C)
}
