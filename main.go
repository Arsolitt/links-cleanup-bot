package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

var (
	ErrNoURL      = errors.New("no URL found")
	ErrInvalidURL = errors.New("invalid URL")
	ErrNotYouTube = errors.New("not a YouTube URL")
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load .env file", "error", err.Error())
		os.Exit(1)
	}
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		slog.Error("TELEGRAM_BOT_TOKEN is not set")
		os.Exit(1)
	}
	logLevel := os.Getenv("LOG_LEVEL")
	var slogLevel slog.Level
	switch logLevel {
	case "debug":
		slogLevel = slog.LevelDebug
	case "info":
		slogLevel = slog.LevelInfo
	case "warn":
		slogLevel = slog.LevelWarn
	default:
		slogLevel = slog.LevelInfo
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slogLevel})))
	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
		bot.WithErrorsHandler(func(err error) {
			if strings.Contains(err.Error(), "context canceled") {
				slog.Info("Bot stopped")
				return
			}
			slog.Error("Bot error", "error", err.Error())
		}),
	}
	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal("Failed to create bot: ", err)
	}
	slog.Info("Bot started")
	b.Start(ctx)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	slog.Debug("New update")
	if update.Message == nil {
		slog.Debug("Message is nil")
		return
	}
	if update.Message.Text == "" {
		slog.Debug("Message text is empty")
		return
	}
	url, err := ExtractURL(update.Message.Text)
	if err != nil {
		slog.Debug("Failed to extract URL", "error", err.Error())
		return
	}
	cleanedURL, err := CleanYouTubeURL(url)
	if err != nil {
		slog.Debug("Failed to clean YouTube URL", "error", err.Error())
		return
	}
	slog.Debug("Sending message", "cleanedURL", cleanedURL, "chatID", update.Message.Chat.ID, "messageID", update.Message.Chat.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   cleanedURL,
		ReplyParameters: &models.ReplyParameters{
			MessageID:                update.Message.ID,
			ChatID:                   update.Message.Chat.ID,
			AllowSendingWithoutReply: true,
		},
	})
}
