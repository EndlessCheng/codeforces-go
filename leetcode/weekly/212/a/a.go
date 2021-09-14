package main

// github.com/EndlessCheng/codeforces-go
func slowestKey(releaseTimes []int, keysPressed string) (ans byte) {
	releaseTimes = append([]int{0}, releaseTimes...)
	mx := 0
	for i := range keysPressed {
		b := keysPressed[i]
		t := releaseTimes[i+1] - releaseTimes[i]
		if t > mx {
			ans = b
			mx = t
		} else if t == mx && b > ans {
			ans = b
		}
	}
	return
}
