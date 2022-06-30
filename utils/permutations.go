package utils

func Permutations(items []string) (result [][]string) {
	if len(items) == 1 {
		return append(result, items)
	}

	for _, p := range Permutations(items[1:]) {
		for i, n := 0, len(p)+1; i < n; i++ {
			perm := make([]string, 0, n)
			perm = append(perm, p[:i]...)
			perm = append(perm, items[0])
			perm = append(perm, p[i:]...)
			result = append(result, perm)
		}
	}
	return
}
