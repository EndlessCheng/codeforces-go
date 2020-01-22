package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	// copy to the Custom Testcase
	const exampleIns = `
[2,1,4]
[1,0,3]
[0,-10,10]
[5,1,7,0,2]
[]
[5,1,7,0,2]
[0,-10,10]
[]
[1,null,8]
[8,1]
`
	exampleOuts := `
[0,1,1,2,3,4]
[-10,0,0,1,2,5,7,10]
[0,1,2,5,7]
[-10,0,10]
[1,1,8,8]
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
[0,1,1,2,3,4]
[-10,0,0,1,2,5,7,10]
[0,1,2,5,7]
[-10,0,10]
[1,1,8,8]
`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
