package main

func collections() {
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

	_ = []interface{}{toBytes}
}
