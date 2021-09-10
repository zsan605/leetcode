package problems

import "strconv"

func addBinary(a string, b string) string {

	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	var carray = 0
	var ret = ""
	for i := 0; i < n; i++ {
		if i < lenA {
			carray += int(a[lenA-i-1] - '0')
		}

		if i < lenB {
			carray += int(b[lenB-i-1] - '0')
		}
		ret = strconv.Itoa(carray%2) + ret
		carray = carray / 2
	}

	if carray >0 {
		ret = "1" + ret
	}
	return ret
}
