// Given a string and an array of integers representing indices, capitalize all letters at the given indices.
//
// For example:
// capitalize("abcdef",[1,2,5]) = "aBCdeF"
// capitalize("abcdef",[1,2,5,100]) = "aBCdeF". There is no index 100.
//
// The input will be a lowercase string with no spaces and an array of digits.
// Good luck!

package kata

import (
	"strings"
)

func Capitalize(st string, arr []int) string {
	result := []rune(st)

	for _, v := range arr {
		if v < len(st) {
			result[v] = rune(strings.ToUpper(string(result[v]))[0])
		}
	}

	return string(result)
}



// Best Practices

package kata

import (
	"unicode"
)

func Capitalize(s string, a []int) string {
  r := []rune(s)
  for _, v := range a {
    if v>=0 && v<len(r) {
      r[v] = unicode.ToUpper(r[v])
    }
  }
  return string(r)
}