package dynamicp

/*
	solution 1: recursive: treat every i as root and unique number is left(i)*right(i)
	solution 2: dp
		G(n) = F(1, n) + F(2, n) + ... + F(n, n).
		F(i, n) = G(i-1)*G(n-i)
		G(n) = G(0)*G(n-1) + G(1)*G(n-2) ... + G(n-1)*G(0)
*/

func numTrees(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}
