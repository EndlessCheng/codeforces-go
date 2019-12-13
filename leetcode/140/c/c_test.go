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
[1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14]
1
[5,4,8,11,null,17,4,7,1,null,null,5,3]
22
[5,-6,-6]
0
`
	exampleOuts := `
[1,2,3,4,null,null,7,8,9,null,14]
[5,4,8,11,null,17,4,7,null,null,null,5]
[]
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
[1,2,3,4,null,null,7,8,9,null,14]
[5,4,8,11,null,17,4,7,null,null,null,5]
[]
`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
