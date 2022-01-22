package array

// fromで渡した配列からtargetの配列を取り除く
//
// @param from
// @param target
// @return result
func Take(from, target []int) []int {
	result := from
	for _, i := range target {
		list := make([]int, 0)
		for _, j := range result {
			if i != j {
				list = append(list, j)
			}
		}
		result = list
	}
	return result
}

// 受け取った配列を逆順に並べ替える
//
// @param src
// @return dst
func Reverse(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	for i := len(dst)/2 - 1; i >= 0; i-- {
		opp := len(dst) - 1 - i
		dst[i], dst[opp] = dst[opp], dst[i]
	}
	return dst
}

// 受け取った配列を指定されたサイズで分割する
//
// @param src
// @param size
// @return dst
func Split(src []int, size int) [][]int {
	dst := make([][]int, 0, (len(src)+size-1)/size)
	for size < len(src) {
		src, dst = src[size:], append(dst, src[0:size:size])
	}
	dst = append(dst, src)
	return dst
}

// 受け取った配列を指定された条件でフィルタリングする
//
// @param src {[]int}
// @param conditions {func(int) bool}
// @return dst {[]int}
func Filter(src []int, conditions func(int) bool) []int {
	dst := make([]int, 0)
	for _, v := range src {
		if conditions(v) {
			dst = append(dst, v)
		}
	}
	return dst
}
