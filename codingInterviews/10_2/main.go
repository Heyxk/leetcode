package main

func numWays(n int) int {
	const mod int = 1e9 + 7

	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}

	p, q, r := 1, 1, 2
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}
