package messenger

import (
	"github.com/twitchyliquid64/CNC/logging"
  "github.com/Syfaro/telegram-bot-api"
)

func TelegramMessageHandler() {
  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  err := gTelegramConnection.UpdatesChan(u)
  if err != nil {
      logging.Error("messenger", err)
			tracking_notifyFault(err)
  } else {
    for update := range gTelegramConnection.Updates {

      if (update.Message.NewChatParticipant.UserName == gTelegramBotUsername) {
        onChatJoined(update.Message)
      } else if (update.Message.LeftChatParticipant.UserName == gTelegramBotUsername) {
        onChatLeft(update.Message)
      }else {
        onMessage(update.Message)
      }
    }
  }
}

func onMessage(msg tgbotapi.Message) {
  logging.Info("messenger", "Message from ", msg.From.UserName, " - text: ", msg.Text)
  sendReply(msg, "Noted.")
  sendSimpleMessage(msg.Chat.ID, "Lol.")
}

func onChatJoined(msg tgbotapi.Message) {
  logging.Info("messenger", "Added to new chat: ", msg.Chat.Title)
}

func onChatLeft(msg tgbotapi.Message) {
  logging.Info("messenger", "Removed from chat: ", msg.Chat.Title)
}

func sendSimpleMessage(chatID int, text string) {
  reply := tgbotapi.NewMessage(chatID, text)
  gTelegramConnection.SendMessage(reply)
}

func sendReply(msg tgbotapi.Message, text string) {
  reply := tgbotapi.NewMessage(msg.Chat.ID, text)
  reply.ReplyToMessageID = msg.MessageID
  gTelegramConnection.SendMessage(reply)
}
