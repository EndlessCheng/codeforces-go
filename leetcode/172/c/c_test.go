package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	// copy to the Custom Testcase
	const exampleIns = `
[1,2,3,2,null,2,4]
2
[1,3,3,3,2]
3
[1,2,null,2,null,2]
2
[1,1,1]
1
[1,2,3]
2
`
	exampleOuts := `
[1,null,3,null,4]
[1,3,null,null,2]
[1]
[]
[1,2,3]
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
[1,null,3,null,4]
[1,3,null,null,2]
[1]
[]
[1,null,3]
`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
