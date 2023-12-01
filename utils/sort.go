package sort

func Max(list []int) int {
	switch len(list) {
	case 0:
		return 0
	default:
		max := list[0]
		for _, val := range list[0:] {
			if max < val {
				max = val
			}
		}
		return max
	}
}

func Sum(list []int) int {
	result := 0
	for _, v := range list {
		result += v
	}
	return result
}

func MergeSort(list []int) []int {
	listLen := len(list)
	switch listLen {
	case 0:
		return nil
	case 1:
		return list
	default:
		halfLen := listLen / 2
		leftSide := list[0:halfLen]
		rightSide := list[halfLen:]

		mergedLeft := MergeSort(leftSide)
		mergedRight := MergeSort(rightSide)

		return merge(mergedLeft, mergedRight)
	}
}

// REQUIREMENT: la and lb are internally-sorted in ascending order
func merge(la, lb []int) []int {
	var mergedList []int

	for ia, ib := 0, 0; ia <= len(la) || ia <= len(lb); {
		if ia == len(la) {
			mergedList = append(mergedList, lb[ib:]...)
			break
		} else if ib == len(lb) {
			mergedList = append(mergedList, la[ia:]...)
			break
		} else {
			if la[ia] > lb[ib] {
				mergedList = append(mergedList, lb[ib])
				ib++
			} else {
				mergedList = append(mergedList, la[ia])
				ia++
			}
		}
	}

	return mergedList // Dead line, needed for compilation
}
