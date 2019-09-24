package copypasta

func commonFunc() {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
}
