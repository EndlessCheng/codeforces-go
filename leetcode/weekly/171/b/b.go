package main

func minFlips(a int, b int, c int) (ans int) {
	for i := uint(0); i < 31; i++ {
		if c>>i&1 == 0 {
			if a>>i&1 == 1 {
				ans++
			}
			if b>>i&1 == 1 {
				ans++
			}
		} else {
			if a>>i&1 == 1 {
				continue
			}
			if b>>i&1 == 1 {
				continue
			}
			ans++
		}
	}
	return
}
