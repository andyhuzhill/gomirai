package gomirai

import (
	"github.com/andyhuzhill/gomirai/bot"
	"github.com/andyhuzhill/gomirai/message"
)

func SendGroupMessageWithBot(b *bot.Bot, qq, quote uint, msg ...message.Message) (uint, error) {
	return b.SendGroupMessage(qq, quote, msg...)
}

func SendFriendMessageWithBot(b *bot.Bot, group, quote uint, msg ...message.Message) (uint, error) {
	return b.SendGroupMessage(group, quote, msg...)
}
