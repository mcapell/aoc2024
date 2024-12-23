package solver

func ReverseNewSlice(slice []int) []int {
	n := slice
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return n
}
