package pkg_logger

import (
	"context"
	"errors"
	"product-mall/internal/constants"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GORMLogger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
	LogLevel              int
}

func NewGORMLogger() *GORMLogger {
	return &GORMLogger{
		SkipErrRecordNotFound: true,
		Debug:                 true, //是否打印debug模式的sql
		LogLevel:              4,
	}
}

func (l *GORMLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *GORMLogger) Info(ctx context.Context, s string, args ...interface{}) {
	GetLoggerWithCtx(ctx).Infof(s, args...)
}

func (l *GORMLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	GetLoggerWithCtx(ctx).Warnf(s, args...)
}

func (l *GORMLogger) Error(ctx context.Context, s string, args ...interface{}) {
	GetLoggerWithCtx(ctx).Errorf(s, args...)
}

func (l *GORMLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		GetLoggerWithCtx(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		GetLoggerWithCtx(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	if l.Debug {
		GetLoggerWithCtx(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
		return
	}
}

func GetLoggerWithCtx(ctx context.Context) *logrus.Entry {
	// 从context中获取request_id
	requestId, ok := ctx.Value(constants.HeaderXRequestID).(string)
	if !ok {
		requestId = ""
	}
	return LogrusObj.WithFields(logrus.Fields{
		constants.HeaderXRequestID: requestId,
	})
}
