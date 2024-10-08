// Complete the solution so that it reverses all of the words within the string passed in.
// Words are separated by exactly one space and there are no leading or trailing spaces.
//
// Example(Input --> Output):
// "The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"

package kata

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	s := ""
	words := strings.Fields(str)
	for i := len(words); i > 0; i-- {
		if i == 1 {
			s += fmt.Sprintf("%s", words[i-1])
		} else {
			s += fmt.Sprintf("%s ", words[i-1])
		}
	}
	return s
}

// Best Practices

package kata

import "strings"

func ReverseWords(str string) string {
  words := strings.Split(str, " ")
  for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
  }
  return strings.Join(words, " ")
}