package cyclic_rotation

func Solution(K int, arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return arr
	}
	k := K % len(arr)
	return append(arr[len(arr)-k:], arr[:len(arr)-k]...)
}
