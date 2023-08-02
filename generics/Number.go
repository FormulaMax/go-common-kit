package generics

type Number interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 |
		~float32 | ~float64 | ~byte
}
