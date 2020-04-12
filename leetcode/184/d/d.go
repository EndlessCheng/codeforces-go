package main

// https://oeis.org/A052913
// a(n) is also the number of 3 X n matrices of integers for which the upper-left hand corner is a 1,
// the rows and columns are weakly increasing, and two adjacent entries differ by at most 1.
//   - Richard Stanley, Jun 06 2010
func numOfWays(n int) (ans int) {
	const mod int = 1e9 + 7
	a := make([]int, n+1)
	a[0] = 3
	a[1] = 12
	for i := 2; i <= n; i++ {
		a[i] = (5*a[i-1] - 2*a[i-2]) % mod
	}
	return (a[n] + mod) % mod
}
