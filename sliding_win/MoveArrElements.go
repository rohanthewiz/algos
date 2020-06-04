package sliding_win

func MoveZerosToArrEnd(arr []int) []int {
	lnArr := len(arr)

	arr2 := make([]int, lnArr)
	i2 := 0

	for i := 0; i < lnArr; i++ {
		if arr[i] != 0 {
			arr2[i2] = arr[i]
			i2++
		}
	}

	return arr2
}

func moveOnesToArrStart(arr []int) []int {
	lnArr := len(arr)
	sIdx := 0
	eIdx := lnArr - 1
	lIdx := eIdx

	arr2 := make([]int, lnArr)

	for i := 0; i < lnArr; i++ {
		if arr[lIdx-i] == 1 {
			arr2[sIdx] = 1
			sIdx++
		} else {
			arr2[eIdx] = arr[lIdx-i]
			eIdx--
		}
	}

	return arr2
}
