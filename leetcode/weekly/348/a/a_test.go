// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, minimizedStringLength, "a.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
	if err := testutil.RunFuncWithRandomInput(t, minimizedStringLength); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-348/problems/minimize-string-length/
// https://leetcode.cn/problems/minimize-string-length/
