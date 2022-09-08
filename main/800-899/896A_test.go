package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/896/A
// https://codeforces.com/problemset/status/896/problem/A
func TestCF896A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 1
1 2
1 111111111111
outputCopy
Wh.
inputCopy
5
0 69
1 194
1 139
0 47
1 66
outputCopy
abdef
inputCopy
10
4 1825
3 75
3 530
4 1829
4 1651
3 187
4 584
4 255
4 774
2 474
outputCopy
Areyoubusy
inputCopy
1
999 1000000000000000000
outputCopy
?`
	testutil.AssertEqualCase(t, rawText, -1, CF896A)
}
