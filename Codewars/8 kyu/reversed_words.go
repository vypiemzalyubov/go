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
  strSplit := strings.Split(str, " ")
  var res []string
  for i := len(strSplit)-1; i >= 0; i--{
   res = append(res, strSplit[i]) 
  }
  return strings.Join(res, " ")
}