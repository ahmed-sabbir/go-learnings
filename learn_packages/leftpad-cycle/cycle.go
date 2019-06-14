package cycle

import (
	"github.ibm.com/Sabbir-Ahmed/learn_packages/leftpad"
)

var defaultChar = ' '

// FormatDouble is useless function
func FormatDouble(s string, i int) string {
	return leftpad.Format(s+s, i)
}
