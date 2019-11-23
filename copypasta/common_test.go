package copypasta

import "testing"

func Test_grayCode(t *testing.T) {
	//t.Log(grayCode(3))
	//t.Log(grayCode(4))

	a := []int{0,1, 6, 19}
	j := 0
	for i := 0; i < 20; i++ {
		if j == len(a) || i < a[j] {
			// do i
			t.Log(i)
		} else {
			j++
		}
	}
}
