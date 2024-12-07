package slices

func Copy[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

func IndexOf[T comparable](src []T, v T) int {
	for i, value := range src {
		if value == v {
			return i
		}
	}

	return -1
}

func DeepCopy2d[T any](src [][]T) [][]T {
	var dst [][]T
	for _, r := range src {
		dst = append(dst, append([]T{}, r...))
	}
	return dst
}
