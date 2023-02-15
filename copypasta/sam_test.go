package copypasta

import "testing"

func Test_sam_printSAM(t *testing.T) {
	s := newSam("aababa")
	s.printSAM()
}
