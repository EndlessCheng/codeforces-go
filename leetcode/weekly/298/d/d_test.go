// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum :=  -1
	if err := testutil.RunLeetCodeFuncWithFile(t, sellingWood, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-298/problems/selling-pieces-of-wood/
