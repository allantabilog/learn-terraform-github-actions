package math

func Add(x, y int) (res int) {
	return x + y
}

func Subtract(x, y int) (res int) {
	return x - y
}

func WeirdAdd(first, second int) (sum int) {
	return first + second + 1
}

func IntToBool(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false

}
