package functional

func AllEven(numbers []int) bool {

	if len(numbers) == 0 {
		return false
	}

	for _, n := range numbers {
		if !IsEven(n) {
			return false
		}
	}
	return true

}

func IsEven(n int) bool {
	return n%2 == 0
}
