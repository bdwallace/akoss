package bot

import (
	"github.com/astaxie/beego"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"library/common"
)

type Telegram struct {
	ToKen  string
	ChatId int64
	Token  string
	Bot    *tgbotapi.BotAPI
}

// 初始化机器人信息
func (c *Telegram) AkReptileInit() (err error) {
	// 机器人的token
	// c.ToKen = "1203023143:AAF0evjZCJtL-oz-7vCFlEarlO78VDczxic"
	c.ToKen = beego.AppConfig.String("telegramToken")
	// 聊天组的chanId
	// c.ChatId = -417018845
	c.ChatId = common.StrToInt64(beego.AppConfig.String("telegramChatId"))

	c.Bot, err = tgbotapi.NewBotAPI(c.ToKen)
	if err != nil {
		beego.Error("init telegram bot err: ", err)
	}

	// 开启会有详细的打印
	// c.Bot.Debug = true

	return
}

func (c *Telegram) SendText(text string) (err error) {

	msg := tgbotapi.NewMessage(c.ChatId, text)
	_, err = c.Bot.Send(msg)
	return
}

func (c *Telegram) SendPhotoFile(photo string) (err error) {
	msg := tgbotapi.NewPhotoUpload(c.ChatId, photo)
	msg.Caption = "Test"
	_, err = c.Bot.Send(msg)
	return
}

func (c *Telegram) SendPhoto(photo []byte, name string) (err error) {
	b := tgbotapi.FileBytes{Name: "image.jpg", Bytes: photo}
	msg := tgbotapi.NewPhotoUpload(c.ChatId, b)
	msg.Caption = name
	_, err = c.Bot.Send(msg)
	return
}
