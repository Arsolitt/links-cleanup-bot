# Links Cleanup Bot

Telegram бот для очистки YouTube ссылок от трекинговых параметров.

## Что делает

Удаляет параметр `si` из YouTube ссылок. Получает сообщение с `youtu.be` или `youtube.com` ссылкой, возвращает очищенную версию.

**До:** `https://youtu.be/dQw4w9WgXcQ?si=abc123def456`  
**После:** `https://youtu.be/dQw4w9WgXcQ`

## Запуск

1. Создай `.env` файл:
```env
TELEGRAM_BOT_TOKEN=your_token_here
LOG_LEVEL=info
```

2. Запусти:
```bash
go run .
```

Или через Docker:
```bash
docker build -t links-cleanup-bot .
docker run --env-file .env links-cleanup-bot
```

## Сборка

```bash
go build -o links-cleanup-bot
```

