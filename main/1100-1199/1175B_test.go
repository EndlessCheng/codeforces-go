package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1175B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
add
for 43
end
for 10
for 15
add
end
add
end
outputCopy
161
inputCopy
2
for 62
end
outputCopy
0
inputCopy
11
for 100
for 100
for 100
for 100
for 100
add
end
end
end
end
end
outputCopy
OVERFLOW!!!`
	testutil.AssertEqualCase(t, rawText, 0, CF1175B)
}
