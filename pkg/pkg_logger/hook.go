package pkg_logger

import (
	"product-mall/internal/constants"

	"github.com/sirupsen/logrus"
)

// hook 增加解析reques-id的
type LogTrace struct {
}

func NewLogTrace() LogTrace {
	return LogTrace{}
}

func (hook LogTrace) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook LogTrace) Fire(entry *logrus.Entry) error {
	ctx := entry.Context
	if ctx != nil {
		traceId, ok := ctx.Value(constants.HeaderXRequestID).(string)
		if ok {
			entry.Data[constants.HeaderXRequestID] = traceId
		}
	}
	return nil
}
