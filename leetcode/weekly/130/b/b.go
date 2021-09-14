package main

func baseNeg2(n int) (res string) {
	if n == 0 {
		return "0"
	}
	for ; n != 0; n = -(n >> 1) {
		res = string('0'+n&1) + res
	}
	return
}
