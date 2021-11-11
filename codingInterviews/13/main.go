package main

type ret struct {
	mm  int
	nn  int
	res int
	a   [100][100]int
}

func movingCount(m int, n int, k int) int {
	r := &ret{mm: m, nn: n, res: 0}
	r.mm, r.nn = m, n
	r.dfs(0, 0, k)
	return r.res

}

func (r *ret) dfs(m int, n int, k int) {
	if m >= r.mm || m < 0 || n >= r.nn || n < 0 || (m/10+m%10+n/10+n%10) > k {
		return
	}
	if r.a[m][n] == 1 {
		return
	}
	r.res++
	r.a[m][n] = 1
	r.dfs(m+1, n, k)
	r.dfs(m-1, n, k)
	r.dfs(m, n+1, k)
	r.dfs(m, n-1, k)
}
