
# slog: Microsoft Teams handler

[![tag](https://img.shields.io/github/tag/samber/slog-microsoft-teams.svg)](https://github.com/samber/slog-microsoft-teams/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/slog-microsoft-teams?status.svg)](https://pkg.go.dev/github.com/samber/slog-microsoft-teams)
![Build Status](https://github.com/samber/slog-microsoft-teams/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/slog-microsoft-teams)](https://goreportcard.com/report/github.com/samber/slog-microsoft-teams)
[![Coverage](https://img.shields.io/codecov/c/github/samber/slog-microsoft-teams)](https://codecov.io/gh/samber/slog-microsoft-teams)
[![Contributors](https://img.shields.io/github/contributors/samber/slog-microsoft-teams)](https://github.com/samber/slog-microsoft-teams/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/slog-microsoft-teams)](./LICENSE)

A [Teams](https://www.microsoft.com/en/microsoft-teams) Handler for [slog](https://pkg.go.dev/log/slog) Go library.

**See also:**

- [slog-multi](https://github.com/samber/slog-multi): `slog.Handler` chaining, fanout, routing, failover, load balancing...
- [slog-formatter](https://github.com/samber/slog-formatter): `slog` attribute formatting
- [slog-sampling](https://github.com/samber/slog-sampling): `slog` sampling policy
- [slog-gin](https://github.com/samber/slog-gin): Gin middleware for `slog` logger
- [slog-echo](https://github.com/samber/slog-echo): Echo middleware for `slog` logger
- [slog-fiber](https://github.com/samber/slog-fiber): Fiber middleware for `slog` logger
- [slog-datadog](https://github.com/samber/slog-datadog): A `slog` handler for `Datadog`
- [slog-rollbar](https://github.com/samber/slog-rollbar): A `slog` handler for `Rollbar`
- [slog-sentry](https://github.com/samber/slog-sentry): A `slog` handler for `Sentry`
- [slog-syslog](https://github.com/samber/slog-syslog): A `slog` handler for `Syslog`
- [slog-logstash](https://github.com/samber/slog-logstash): A `slog` handler for `Logstash`
- [slog-fluentd](https://github.com/samber/slog-fluentd): A `slog` handler for `Fluentd`
- [slog-graylog](https://github.com/samber/slog-graylog): A `slog` handler for `Graylog`
- [slog-loki](https://github.com/samber/slog-loki): A `slog` handler for `Loki`
- [slog-slack](https://github.com/samber/slog-slack): A `slog` handler for `Slack`
- [slog-telegram](https://github.com/samber/slog-telegram): A `slog` handler for `Telegram`
- [slog-mattermost](https://github.com/samber/slog-mattermost): A `slog` handler for `Mattermost`
- [slog-microsoft-teams](https://github.com/samber/slog-microsoft-teams): A `slog` handler for `Microsoft Teams`
- [slog-webhook](https://github.com/samber/slog-webhook): A `slog` handler for `Webhook`
- [slog-kafka](https://github.com/samber/slog-kafka): A `slog` handler for `Kafka`

## 🚀 Install

```sh
go get github.com/samber/slog-microsoft-teams
```

**Compatibility**: go >= 1.21

No breaking changes will be made to exported APIs before v2.0.0.

## 💡 Usage

GoDoc: [https://pkg.go.dev/github.com/samber/slog-microsoft-teams](https://pkg.go.dev/github.com/samber/slog-microsoft-teams)

### Handler options

```go
type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// Teams webhook url
	WebhookURL string

	// optional: customize Teams event builder
	Converter Converter
}
```

### Example

```go
import (
	slogmicrosoftteams "github.com/samber/slog-microsoft-teams"
	"log/slog"
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

```

## 🤝 Contributing

- Ping me on twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/slog-microsoft-teams)
- Fix [open issues](https://github.com/samber/slog-microsoft-teams/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=samber/slog-microsoft-teams)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2023 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
