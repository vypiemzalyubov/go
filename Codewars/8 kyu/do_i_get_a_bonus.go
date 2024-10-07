// It's bonus time in the big city! The fatcats are rubbing their paws in anticipation... but who is going to make the most money?
// Build a function that takes in two arguments (salary, bonus). Salary will be an integer, and bonus a boolean.
// If bonus is true, the salary should be multiplied by 10. If bonus is false, the fatcat did not make enough money and must receive only his stated salary.
// Return the total figure the individual will receive as a string prefixed with "£" (= "\u00A3", JS, Go, Java, Scala, and Julia), "$" (C#, C++, Dart, Ruby, Clojure, Elixir, PHP, Python, Haskell, and Lua) or "¥" (Rust).

package kata

import (
	"fmt"
)

func BonusTime(salary int, bonus bool) string {
	if bonus {
		return fmt.Sprintf("£%v", salary*10)
	}
	return fmt.Sprintf("£%v", salary)
}



// Best Practices

package kata

import "strconv"

func BonusTime(salary int, bonus bool) string {
  if bonus {
    salary *= 10
  }
  
  return "£" + strconv.Itoa(salary) 
}