package slogmicrosoftteams

import (
	"fmt"

	"log/slog"

	slogcommon "github.com/samber/slog-common"
)

var SourceKey = "source"

type Converter func(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) string

func DefaultConverter(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) string {
	// aggregate all attributes
	attrs := slogcommon.AppendRecordAttrsToAttrs(loggerAttr, groups, record)

	// developer formatters
	if addSource {
		attrs = append(attrs, slogcommon.Source(SourceKey, record))
	}
	attrs = slogcommon.ReplaceAttrs(replaceAttr, []string{}, attrs...)
	attrs = slogcommon.RemoveEmptyAttrs(attrs)

	// handler formatter
	message := attrToTeamsMessage("", attrs)
	return message
}

func attrToTeamsMessage(base string, attrs []slog.Attr) string {
	message := ""

	for i := range attrs {
		attr := attrs[i]
		k := base + attr.Key
		v := attr.Value
		kind := attr.Value.Kind()

		if kind == slog.KindGroup {
			message += attrToTeamsMessage(k+".", v.Group())
		} else {
			message += fmt.Sprintf("**%s**: %s\n\n", k, slogcommon.ValueToString(v))
		}
	}

	return message
}
