package go_common_kit

import "github.com/FormulaMax/go-common-kit/generics"

type Comparator[T any] func(src T, dst T) int

func ComparatorRealNumber[T generics.RealNumber](src, dst T) int {
	switch {
	case src < dst:
		return -1
	case src == dst:
		return 0
	default:
		return 1
	}
}
