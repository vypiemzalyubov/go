// There is an array with some numbers. All numbers are equal except for one. Try to find it!
//
// findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
// findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
//
// It’s guaranteed that array contains at least 3 numbers.
// The tests contain some very huge arrays, so think about performance.

package kata

import "sort"

func FindUniq(arr []float32) float32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[j] < arr[i]
	})
	switch {
	case arr[0] == arr[1]:
		return arr[len(arr)-1]
	default:
		return arr[0]
	}
}



// Best Practices

package kata

func FindUniq(arr []float32) float32 {
  if arr[0] != arr[1] && arr[1] == arr[2] { return arr[0] }
  for i, v := range arr[1:] {
    if v != arr[i] { return v }
  }
  return 0
}