package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(errors validate.Errors) string {
	res := ""
	for _, val := range errors.Errors {
		res += strings.Join(val, ", ") + "\n"
	}
	return res
}
