package slogmicrosoftteams

import (
	"context"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
	slogcommon "github.com/samber/slog-common"

	"log/slog"
)

type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// Teams webhook url
	WebhookURL string

	// optional: customize Teams event builder
	Converter Converter

	// optional: see slog.HandlerOptions
	AddSource   bool
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

func (o Option) NewMicrosoftTeamsHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.WebhookURL == "" {
		panic("missing Teams webhook url")
	}

	return &MicrosoftTeamsHandler{
		option: o,
		attrs:  []slog.Attr{},
		groups: []string{},
	}
}

var _ slog.Handler = (*MicrosoftTeamsHandler)(nil)

type MicrosoftTeamsHandler struct {
	option Option
	attrs  []slog.Attr
	groups []string
}

func (h *MicrosoftTeamsHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *MicrosoftTeamsHandler) Handle(ctx context.Context, record slog.Record) error {
	converter := DefaultConverter
	if h.option.Converter != nil {
		converter = h.option.Converter
	}

	message := converter(h.option.AddSource, h.option.ReplaceAttr, h.attrs, h.groups, &record)

	mstClient := goteamsnotify.NewTeamsClient()

	msgCard := messagecard.NewMessageCard()
	msgCard.Title = record.Message
	msgCard.Text = message
	msgCard.ThemeColor = ColorMapping[record.Level]

	go func() {
		_ = mstClient.Send(h.option.WebhookURL, msgCard)
	}()

	return nil
}

func (h *MicrosoftTeamsHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &MicrosoftTeamsHandler{
		option: h.option,
		attrs:  slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...),
		groups: h.groups,
	}
}

func (h *MicrosoftTeamsHandler) WithGroup(name string) slog.Handler {
	return &MicrosoftTeamsHandler{
		option: h.option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}
