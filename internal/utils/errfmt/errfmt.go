package errfmt

import (
	"github.com/enginer/jet/v2/internal/utils/is"
	"strings"
)

// Trace returns well formatted wrapped error trace string
func Trace(err error) string {
	if is.Nil(err) {
		return ""
	}
	return "Error trace:\n" + " - " + strings.Replace(err.Error(), ": ", ":\n - ", -1)
}
