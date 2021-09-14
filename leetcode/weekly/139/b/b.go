package main

func maxEqualRowsAfterFlips_N3(mat [][]int) (ans int) {
	n, m := len(mat), len(mat[0])
	flipCols := func(flips []bool) {
		for j, f := range flips {
			if f {
				for i := range mat {
					mat[i][j] ^= 1
				}
			}
		}
	}
	seen := make([][2]bool, n)
	for i, row := range mat {
		// change all elements of this row 0 or 1
		for v := 0; v <= 1; v++ {
			if seen[i][v] {
				continue
			}
			flips := make([]bool, m)
			for j, mij := range row {
				flips[j] = mij != v
			}
			flipCols(flips)
			cnt := 1
			for j := i + 1; j < n; j++ {
				rowJ := mat[j]
				allSame := true
				for _, mij := range rowJ {
					if mij != rowJ[0] {
						allSame = false
						break
					}
				}
				if allSame {
					cnt++
					seen[j][rowJ[0]] = true
				}
			}
			if cnt > ans {
				ans = cnt
			}
			flipCols(flips)
		}
	}
	return
}

func maxEqualRowsAfterFlips(mat [][]int) (ans int) {
	m := len(mat[0])
	mp := map[string]int{}
	s1, s2 := make([]byte, m), make([]byte, m)
	for _, row := range mat {
		for j, mij := range row {
			s1[j] = byte(mij)
			s2[j] = byte(mij) ^ 1
		}
		mp[string(s1)]++
		mp[string(s2)]++
	}
	for _, v := range mp {
		if v > ans {
			ans = v
		}
	}
	return
}
