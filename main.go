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
	Token := "7649681669:AAFic1YnN2AP7F7rX9QLu9kwOZ32hqg5uJk"

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
			return // Игнорируем обновления без сообщений
		}

		chatID := tu.ID(update.Message.Chat.ID)

		// Обработка нажатий кнопок или ввода текста
		text := "Бот еще в разработке.." // Default message
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
	}, th.Any()) // Обрабатываем любые сообщения

	bh.Start()
}
