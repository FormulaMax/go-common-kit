package mapx

func Keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

func Values[K comparable, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for k := range m {
		res = append(res, m[k])
	}
	return res
}

func KeysValues[K comparable, V any](m map[K]V) ([]K, []V) {
	resK := make([]K, 0, len(m))
	resV := make([]V, 0, len(m))
	for k := range m {
		resK = append(resK, k)
		resV = append(resV, m[k])
	}
	return resK, resV
}
