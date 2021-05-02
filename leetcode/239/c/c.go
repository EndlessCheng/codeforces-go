package main

// github.com/EndlessCheng/codeforces-go
func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func nextPermutation(a []byte) {
	n := len(a)
	i := n - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	j := n - 1
	for j >= 0 && a[i] >= a[j] {
		j--
	}
	a[i], a[j] = a[j], a[i]
	reverse(a[i+1:])
}

func mapPos(a string, b []byte) []int {
	pos := [10][]int{}
	for i, v := range a {
		pos[v&15] = append(pos[v&15], i)
	}
	ids := make([]int, len(b))
	for i, v := range b {
		v &= 15
		ids[i] = pos[v][0]
		pos[v] = pos[v][1:]
	}
	return ids
}

func mergeCount(a []int) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	left := append([]int(nil), a[:n/2]...)
	right := append([]int(nil), a[n/2:]...)
	cnt := mergeCount(left) + mergeCount(right)
	l, r := 0, 0
	for i := range a {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			a[i] = left[l]
			l++
		} else {
			cnt += n/2 - l
			a[i] = right[r]
			r++
		}
	}
	return cnt
}

func getMinSwaps(s string, k int) int {
	t := []byte(s)
	for ; k > 0; k-- {
		nextPermutation(t)
	}
	return mergeCount(mapPos(s, t))
}
