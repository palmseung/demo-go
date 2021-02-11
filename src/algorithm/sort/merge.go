package main

//merge sort는 partially sorting된 경우라도 속도는 같다.
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])

	i, j := 0, 0
	rst := make([]int, 0)

	for i < len(left) || j < len(right) {
		if i >= len(left) {
			rst = append(rst, right[j])
			j++
		} else if j >= len(right) {
			rst = append(rst, left[i])
			i++
		} else {
			if left[i] < right[j] {
				rst = append(rst, left[i])
				i++
			} else {
				rst = append(rst, right[j])
				j++
			}
		}

	}
	return rst
}
