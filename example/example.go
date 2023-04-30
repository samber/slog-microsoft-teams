package main

import (
	"fmt"
	"time"

	slogmicrosoftteams "github.com/samber/slog-microsoft-teams"

	"golang.org/x/exp/slog"
)

func main() {
	url := "https://xxxxxx.webhook.office.com/webhookb2/xxxxx@xxxxx/IncomingWebhook/xxxxx/xxxxx"

	logger := slog.New(slogmicrosoftteams.Option{Level: slog.LevelDebug, WebhookURL: url}.NewMicrosoftTeamsHandler())
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")
}
