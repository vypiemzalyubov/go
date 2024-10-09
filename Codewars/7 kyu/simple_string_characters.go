// In this Kata, you will be given a string and your task will be to return a list of ints detailing the count of uppercase letters, lowercase, numbers and special characters (everything else), as follows.
// The order is: uppercase letters, lowercase letters, numbers and special characters.
//
// "*'&ABCDabcde12345" --> [ 4, 5, 5, 3 ]
// More examples in the test cases.
//
// Good luck!

package kata

func Solve(s string) []int {
	bigLetter, smallLetter, digit, specSymbol := 0, 0, 0, 0
	result := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			bigLetter++
		} else if s[i] >= 97 && s[i] <= 122 {
			smallLetter++
		} else if s[i] >= 48 && s[i] <= 57 {
			digit++
		} else {
			specSymbol++
		}
	}

	result = append(result, bigLetter)
	result = append(result, smallLetter)
	result = append(result, digit)
	result = append(result, specSymbol)

	return result
}



// Best Practices

package kata

func Solve(s string) []int {
  r := make([]int, 4)
  for _, c := range s {
    switch {
      case c >= 'A' && c <= 'Z': r[0]++;
      case c >= 'a' && c <= 'z': r[1]++;
      case c >= '0' && c <= '9': r[2]++;
      default: r[3]++;
    }
  }
  return r
}