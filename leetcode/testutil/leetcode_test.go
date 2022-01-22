package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_isTLE(t *testing.T) {
	DebugTLE = 0
	assert.False(t, isTLE(func() {}))

	DebugTLE = time.Second
	assert.False(t, isTLE(func() {}))
	assert.True(t, isTLE(func() { select {} }))
}

func TestRunLeetCodeFunc(t *testing.T) {
	specialF := func(a []string) []string {
		fmt.Println("args:", a)
		return a
	}
	// , in string slice
	data := [][]string{{`["alice,20,800,mtv","alice,50,100,beijing"]`}}
	if err := RunLeetCodeFunc(t, specialF, data, data); err != nil {
		t.Error(err)
	}

	baseF := func(a string, b int, c float64, d bool, e byte) (string, int, float64, bool, byte) {
		fmt.Println("args:", a, b, c, d, e)
		return a, b, c, d, e
	}
	data = [][]string{{`"ac"`, `-123`, `1.23000`, `true`, `"a"`}}
	if err := RunLeetCodeFunc(t, baseF, data, data); err != nil {
		t.Error(err)
	}

	sliceF := func(a []string, b []int, c []float64, d []bool) ([]string, []int, []float64, []bool) {
		fmt.Println("args:", a, b, c, d)
		return a, b, c, d
	}
	data = [][]string{{`["ac","wa","tle"]`, `[-123,3,0,1]`, `[1.23000,3.00000]`, `[true,false,true]`}}
	if err := RunLeetCodeFunc(t, sliceF, data, data); err != nil {
		t.Error(err)
	}
	data = [][]string{{`[]`, `[]`, `[]`, `[]`}}
	if err := RunLeetCodeFunc(t, sliceF, data, data); err != nil {
		t.Error(err)
	}

	matrixF := func(a [][]string, b [][]int) ([][]string, [][]int) {
		fmt.Println("args:", a, b)
		return a, b
	}
	data = [][]string{{`[["ac","wa","tle"],["1"]]`, `[[-123,3,0,1],[1]]`}}
	if err := RunLeetCodeFunc(t, matrixF, data, data); err != nil {
		t.Error(err)
	}
	data = [][]string{{`[["ac"]]`, `[[-123]]`}}
	if err := RunLeetCodeFunc(t, matrixF, data, data); err != nil {
		t.Error(err)
	}
	data = [][]string{{`[[],[],[],[]]`, `[[],[]]`}}
	if err := RunLeetCodeFunc(t, matrixF, data, data); err != nil {
		t.Error(err)
	}
	data = [][]string{{`[[]]`, `[]`}}
	if err := RunLeetCodeFunc(t, matrixF, data, data); err != nil {
		t.Error(err)
	}

	treeNodeF := func(root *TreeNode) *TreeNode {
		fmt.Println(root.toRawString())
		return root
	}
	data = [][]string{{`[]`}, {`[1]`}, {`[1,2]`}, {`[1,2,3]`}, {`[1,null,2,3]`}, {`[5,4,7,3,null,2,null,-1,null,9]`}, {`[1,null,2,3,4,null,null,5,6]`}}
	if err := RunLeetCodeFunc(t, treeNodeF, data, data); err != nil {
		t.Error(err)
	}

	listNodeF := func(head *ListNode) *ListNode {
		fmt.Println(head.toRawString())
		return head
	}
	data = [][]string{{`[]`}, {`[1]`}, {`[1,2,3,4,5]`}}
	if err := RunLeetCodeFunc(t, listNodeF, data, data); err != nil {
		t.Error(err)
	}
}

type foo struct {
}

func constructor() foo {
	return foo{}
}

func (f *foo) F0(a, b int) {
	fmt.Println("f0", a, b)
}

func (f *foo) F1(a, b int) int {
	fmt.Println("f1", a, b)
	return a + b
}

func (f *foo) F2(a, b int) []int {
	fmt.Println("f2", a, b)
	return []int{a, b}
}

func TestRunLeetCodeClass(t *testing.T) {
	sampleIns := []string{`
["foo","f0","f1","f2"]
[[],[10,100],[1,2],[11,22]]
`, `
["foo","f1","f2"]
[[],[-1,-2],[-11,-22]]
`}
	sampleOuts := []string{`[null,null,3,[11,22]]`, `[null,-3,[-11,-22]]`}
	if err := RunLeetCodeClass(t, constructor, sampleIns, sampleOuts); err != nil {
		t.Error(err)
	}
}
