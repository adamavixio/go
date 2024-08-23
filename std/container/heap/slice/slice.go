package slice

func Up[H ~[]E, E any](h H, less func(H, int, int) bool, j int) {
	for {
		i := (j - 1) / 2
		if i == j || !less(h, j, i) {
			break
		}
		Swap(h, i, j)
		j = i
	}
}

func Down[H ~[]E, E any](h H, less func(H, int, int) bool, n, i int) {
	for {
		j := 2*i + 1
		if j >= n || j < 0 {
			break
		}
		if k := j + 1; k < n && less(h, k, j) {
			j = k
		}
		if !less(h, j, i) {
			break
		}
		Swap(h, i, j)
		i = j
	}
}

func Swap[H ~[]E, E any](h H, i, j int) {
	h[i], h[j] = h[j], h[i]
}
