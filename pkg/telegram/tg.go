package telegram

import (
	"PharmacyBot/pkg/webScrapper"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

const TgApiKey = "5606010016:AAHuAM7bBpugCC0hpRflEz_2_VkmWFHKsJY"

type Bot interface {
}

func BotStart() {
	bot, err := tgbotapi.NewBotAPI(TgApiKey)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		searchLink := "https://www.rlsnet.ru/search_result.htm?word=" + update.Message.Text
		results := webScrapper.SearchResults(searchLink)
		text := "*Результаты поиска:*\n"

		if len(results) == 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"*По вашему запросу ничего не найдено, проверьте корректность введенных данных!*")
			msg.ParseMode = tgbotapi.ModeMarkdown
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		} else {
			for key, elem := range results {
				text += "[" + key + "]" + "(" + elem[0] + ")\n"
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ParseMode = tgbotapi.ModeMarkdown
			msg.DisableWebPagePreview = true
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		}

	}
}
