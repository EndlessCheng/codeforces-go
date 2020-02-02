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
[1,2,3,4,5,6]
[1,null,2,3,4,null,null,5,6]
[2,3,9,10,7,8,6,5,4,11,1]
[1,1]
`
	exampleOuts := `
110
90
1025
1
`
	// copy Your answer in the Run Code Result
	yourAnswers := `

`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
