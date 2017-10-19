package fibonacci

var (
	results map[int]int
)

func InitDynamic() {
	results = make(map[int]int)
	results[0] = 0
	results[1] = 1
}

func Dynamic(n int) int {
	if val, ok := results[n]; ok {
		return val
	} else {
		res := Dynamic(n-1) + Dynamic(n-2)
		results[n] = res
		return res
	}
}
