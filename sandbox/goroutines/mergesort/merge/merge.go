package merge

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	done := make(chan bool)
	mid := len(data) / 2

	var left []int

	go func() {
		left = MergeSort(data[:mid])
		done <- true
	}()

	right := MergeSort(data[mid:])
	<-done
	return Merge(left, right)
}
