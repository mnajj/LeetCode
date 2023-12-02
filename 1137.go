func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	f := make([]int, n+1)
	f[1], f[2] = 1, 1
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2] + f[i-3]
	}
	return f[n]
}
