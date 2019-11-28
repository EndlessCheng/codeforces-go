package testutil

import (
	"fmt"
	"testing"
)

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
	data = [][]string{{`"ac"`, `-123`, `1.23`, `true`, `"a"`}}
	if err := RunLeetCodeFunc(t, baseF, data, data); err != nil {
		t.Error(err)
	}

	sliceF := func(a []string, b []int, c []float64, d []bool) ([]string, []int, []float64, []bool) {
		fmt.Println("args:", a, b, c, d)
		return a, b, c, d
	}
	data = [][]string{{`["ac","wa","tle"]`, `[-123,3,0,1]`, `[1.23,3]`, `[true,false,true]`}}
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
}
