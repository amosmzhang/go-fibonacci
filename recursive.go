package fibonacci

func Recursive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Recursive(n-1) + Recursive(n-2)
	}
}
