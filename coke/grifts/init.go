package grifts

import (
	"github.com/BorysDrozhak/buff_start/coke/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
