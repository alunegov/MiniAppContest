package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

type PreCheckoutQuery struct {
	Response handlers.Response
}

func NewPreCheckoutQuery(r handlers.Response) PreCheckoutQuery {
	return PreCheckoutQuery{
		Response: r,
	}
}

func (m PreCheckoutQuery) CheckUpdate(b *gotgbot.Bot, ctx *ext.Context) bool {
	return ctx.PreCheckoutQuery != nil
}

func (m PreCheckoutQuery) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	return m.Response(b, ctx)
}

func (m PreCheckoutQuery) Name() string {
	return fmt.Sprintf("precheckoutquery_%p", m.Response)
}
