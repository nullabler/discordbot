package context

import (
	"context"

	"github.com/unixoff/discord-bot/config"
)

type typeConst string

const (
	CONFIG_COMPONENT_KEY typeConst = "config.component"
)

type Context struct {
	ctx context.Context
}

func New() *Context {
	return &Context{
		ctx: context.Background(),
	}
}

func (c *Context) SetConfig(config *config.Config) {
	c.ctx = context.WithValue(c.ctx, CONFIG_COMPONENT_KEY, config)
}

func (c *Context) Config() *config.Config {
	return c.ctx.Value(CONFIG_COMPONENT_KEY).(*config.Config)
}
