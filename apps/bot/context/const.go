package context

import (
	"context"

	"github.com/unixoff/discord-bot/config"
)

const (
	CONFIG_COMPONENT_KEY = "config.component"
)

type Context struct {
	ctx context.Context
}

func New() *Context {
	return &Context{
		ctx: context.Background(),
	}
}

func (self *Context) SetConfig(config *config.Config) {
	self.ctx = context.WithValue(self.ctx, CONFIG_COMPONENT_KEY, config)
}

func (self *Context) Config() *config.Config {
	return self.ctx.Value(CONFIG_COMPONENT_KEY).(*config.Config)
}
