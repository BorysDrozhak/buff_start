package grifts

import (
	"BorysDrozhak/buffalo_test/coke/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
