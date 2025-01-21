package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	Token := ""

	bot, err := telego.NewBot(Token, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil {
			return 
		}

		chatID := tu.ID(update.Message.Chat.ID)

		
		text := "Бот еще в разработке.." 
		switch strings.ToLower(update.Message.Text) {
		case "запуск нейросети":
			text = "Запускаем нейросеть..."
		case "поддержка":
			text = "Свяжитесь с нами через @keyboard5"
		case "помощь":
			text = "Список доступных команд: /start"
		case "привет":
			text = "Привет, запусти нейросеть."
		case "/start":
			text = "Приветствую, я Чат-Бот со встроенной нейросеть GPT, выбери кнопку из чата, чтобы активировать сценарий!"
		}

		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Запуск нейросети"),
				tu.KeyboardButton("Поддержка"),
				tu.KeyboardButton("Помощь"),
			),
		)

		message := tu.Message(chatID, text).WithReplyMarkup(keyboard)

		_, err := bot.SendMessage(message)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}, th.Any()) 

	bh.Start()
}
