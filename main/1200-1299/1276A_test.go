package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1276/A
// https://codeforces.com/problemset/status/1276/problem/A
func TestCF1276A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
onetwone
testme
oneoneone
twotwo
outputCopy
2
6 3
0

3
4 1 7 
2
1 4
inputCopy
10
onetwonetwooneooonetwooo
two
one
twooooo
ttttwo
ttwwoo
ooone
onnne
oneeeee
oneeeeeeetwooooo
outputCopy
6
18 11 12 1 6 21 
1
1 
1
3 
1
2 
1
6 
0

1
4 
0

1
1 
2
1 11 `
	testutil.AssertEqualCase(t, rawText, 0, CF1276A)
}
