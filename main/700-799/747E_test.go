package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF747E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
hello,2,ok,0,bye,0,test,0,one,1,two,2,a,0,b,0
outputCopy
3
hello test one 
ok bye two 
a b
inputCopy
a,5,A,0,a,0,A,0,a,0,A,0
outputCopy
2
a 
A a A a A
inputCopy
A,3,B,2,C,0,D,1,E,0,F,1,G,0,H,1,I,1,J,0,K,1,L,0,M,2,N,0,O,1,P,0
outputCopy
4
A K M 
B F H L N O 
C D G I P 
E J `
	testutil.AssertEqualCase(t, rawText, 0, CF747E)
}
