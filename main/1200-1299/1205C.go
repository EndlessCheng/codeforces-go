package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1205C() {
	q := func(a, b, c, d int) (pal byte) {
		Println("?", a+1, b+1, c+1, d+1)
		Scan(&pal)
		return pal ^ 1
	}

	var n, t int
	Scan(&n)
	a := make([][]byte, n)
	for i := range a {
		a[i] = make([]byte, n)
	}
	a[0][0] = '1'
	a[0][1] = '0'
	for i := range a {
		if i == 1 {
			a[1][1] = a[0][0] ^ q(0, 0, 1, 1)
			a[1][2] = a[0][1] ^ q(0, 1, 1, 2)
			a[1][0] = a[1][2] ^ q(1, 0, 1, 2)
		} else if i > 1 {
			a[i][0] = a[i-2][0] ^ q(i-2, 0, i, 0)
			a[i][1] = a[i-2][1] ^ q(i-2, 1, i, 1)
		}
		for j := 2; j < n; j++ {
			a[i][j] = a[i][j-2] ^ q(i, j-2, i, j)
		}
		if i > 1 && a[i-2][i-2] != a[i][i] {
			t = i - 1
		}
	}

	pre, v := a[t][t-1], a[t][t]
	tar := v
	var add byte
	if v == a[t-1][t-1] == (pre == a[t][t+1]) {
		add = q(t-1, t-1, t, t+1)
	} else {
		tar = a[t+1][t+1]
		add = q(t, t-1, t+1, t+1)
	}
	rev := pre == tar == (add > 0)

	Println("!")
	for i, row := range a {
		if rev {
			for j := i%2 ^ 1; j < n; j += 2 {
				row[j] ^= 1
			}
		}
		Printf("%s\n", row)
	}
}

//func main() { cf1205C() }
