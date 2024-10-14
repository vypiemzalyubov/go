// You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N.
// Write a method that takes the array as an argument and returns this "outlier" N.
//
// Examples
// [2, 4, 0, 100, 4, 11, 2602, 36] -->  11 (the only odd number)
// [160, 3, 1719, 19, 11, 13, -21] --> 160 (the only even number)

package kata

func FindOutlier(integers []int) int {
	var lastEven, lastOdd int
	evenCount := 0
	oddCount := 0

	for _, num := range integers {
		if num%2 == 0 {
			evenCount++
			lastEven = num
			if evenCount > 1 && oddCount > 0 {
				break
			}
		} else {
			oddCount++
			lastOdd = num
			if oddCount > 1 && evenCount > 0 {
				break
			}
		}
	}

	if evenCount == 1 {
		return lastEven
	}
	return lastOdd
}
