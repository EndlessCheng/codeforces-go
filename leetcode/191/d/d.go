package main

func getProbability(a []int) (ans float64) {
	const mx = 6
	C := [mx + 1][mx + 1]int{}
	for i := 0; i <= mx; i++ {
		C[i][0], C[i][i] = 1, 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	n := len(a)
	sum := 0
	for _, v := range a {
		sum += v
	}
	sum /= 2

	okWays, tot := 0, 0
	var f func(p, s, cntL, cntR, ways int)
	f = func(p, s, cntL, cntR, ways int) {
		if p == n {
			if s == sum {
				if cntL == cntR {
					okWays += ways
				}
				tot += ways
			}
			return
		}
		for i := 0; i <= a[p] && s+i <= sum; i++ {
			cl, cr := cntL, cntR
			if i > 0 {
				cl++
			}
			if i < a[p] {
				cr++
			}
			f(p+1, s+i, cl, cr, ways*C[a[p]][i])
		}
	}
	f(0, 0, 0, 0, 1)
	return float64(okWays) / float64(tot)
}
