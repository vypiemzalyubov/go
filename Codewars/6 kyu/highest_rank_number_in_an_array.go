// Complete the method which returns the number which is most frequent in the given input array. If there is a tie for most frequent number, return the largest number among them.
//
// Note: no empty arrays will be given.
//
// Examples
// [12, 10, 8, 12, 7, 6, 4, 10, 12]              -->  12
// [12, 10, 8, 12, 7, 6, 4, 10, 12, 10]          -->  12
// [12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]  -->   3

package kata

func HighestRank(nums []int) int {
	var maxK, maxV int
	m := make(map[int]int)

	for _, v := range nums {
		m[v] += 1
	}

	for k, v := range m {
		if v > maxV {
			maxV = v
			maxK = k
		} else if v == maxV && k > maxK {
			maxK = k
		}
	}

	return maxK
}



// Best Practices

package kata

func HighestRank(nums []int) int {
  mii, maxK, maxV := map[int]int{}, 0, 0
  for _, v := range nums {
    mii[v]++
    if mii[v] > maxV || (mii[v] == maxV && v > maxK) {
      maxK = v
      maxV = mii[v]
    }
  }
  