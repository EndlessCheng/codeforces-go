package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func CF1155E() {
	const mod = 1e6 + 3
	inv := [mod]int64{0, 1}
	for i := 2; i < mod; i++ {
		inv[i] = int64(mod-mod/i) * inv[mod%i] % mod
	}

	f := [11]int64{}
	for i := range f {
		Println("?", i)
		Scan(&f[i])
		if f[i] == 0 {
			Println("!", i)
			return
		}
		for j := range f {
			if j != i {
				f[i] = f[i] * inv[(i-j+mod)%mod] % mod
			}
		}
	}
	for k := 11; k < mod; k++ {
		fk := int64(0)
		mul := int64(1)
		for j := range f {
			mul = mul * int64(k-j) % mod
		}
		for i, fi := range f {
			fk = (fk + fi*mul%mod*inv[k-i]) % mod
		}
		if fk == 0 {
			Println("!", k)
			return
		}
	}
	Println("! -1")
}

//func main() { CF1155E() }
