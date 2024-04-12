package go_common_kit

func ToPtr[T any](t T) *T {
	return &t
}
