package main

func collections() {
	var mat [1][1]int

	n, m := len(mat), len(mat[0])

	dir4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	toBytes := func(g [][]string) [][]byte {
		n, m := len(g), len(g[0])
		bytes := make([][]byte, n)
		for i := range bytes {
			bytes[i] = make([]byte, m)
			for j := range bytes[i] {
				bytes[i][j] = g[i][j][0]
			}
		}
		return bytes
	}

	_ = []interface{}{n, m, dir4, toBytes}
}
