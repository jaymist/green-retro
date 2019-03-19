package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jaymist/greenretro/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
