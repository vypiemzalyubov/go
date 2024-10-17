// Write a simple parser that will parse and run Deadfish.
//
// Deadfish has 4 commands, each 1 character long:
// - i increments the value (initially 0)
// - d decrements the value
// - s squares the value
// - o outputs the value into the return array
//
// Invalid characters should be ignored.
//
// Parse("iiisdoso") == []int{8, 64}

package kata

func Parse(data string) []int {
	result := make([]int, 0)
	cnt := 0

	for _, v := range data {
		switch string(v) {
		case "i":
			cnt++
		case "d":
			cnt--
		case "s":
			cnt *= cnt
		case "o":
			result = append(result, cnt)
		default:
			continue
		}
	}

	return result

}
