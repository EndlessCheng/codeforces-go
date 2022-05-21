package copypasta

import "testing"

func Test_sam_printSAM(t *testing.T) {
	s := newSam()
	s.buildSam("aababa")
	s.printSAM()
}
