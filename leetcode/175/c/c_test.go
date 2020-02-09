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
["TweetCounts","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","getTweetCountsPerFrequency","recordTweet","getTweetCountsPerFrequency"]
[[],["tweet3",0],["tweet3",60],["tweet3",10],["minute","tweet3",0,59],["minute","tweet3",0,60],["tweet3",120],["hour","tweet3",0,210]]
`
	exampleOuts := `
[null,null,null,null,[2],[2,1],null,[4]]
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
[null,null,null,null,[2],[2,1],null,[4]]
`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
