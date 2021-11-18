package problems

/*
合并规则numA， numB

numA[left] >= numB[left] or numA[left] <= numB[right]
numA[right] >= numB[left] and numA[left] <= numB[left] and numA[right] >= numB[right], numA
numA[right] >= numB[left] and numA[left] <= numB[left] and numA[right] < numB[right], num[A.left,B.right]
numA[right] >= numB[left] and numA[left] > numB[left] and numA[right] >= numB[right], num[B.left,A.right]
numA[right] >= numB[left] and numA[left] > numB[left] and numA[right] < numB[right], numB
numB[right] >= numA[left]
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	
	for i, val := range intervals {
		val.
	}
}
