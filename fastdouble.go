package fibonacci

func FastDouble(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else if n%2 == 0 {
		half := FastDouble(n / 2)
		next := FastDouble(n/2 + 1)
		return half * (2*next - half)
	} else {
		index := (n - 1) / 2
		half := FastDouble(index)
		next := FastDouble(index + 1)
		return next*next + half*half
	}
}
