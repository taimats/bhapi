package loggers

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewRequestLoggerConfig(logger *slog.Logger) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogRemoteIP: true,
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogError:    true,
		LogLatency:  true,

		HandleError: true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("remote_ip", v.RemoteIP),
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
					slog.String("latency", v.Latency.String()),
				)
				return nil
			}
			logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
				slog.String("remote_ip", v.RemoteIP),
				slog.String("method", v.Method),
				slog.String("uri", v.URI),
				slog.Int("status", v.Status),
				slog.String("latency", v.Latency.String()),
			)
			return nil
		},
	}
}

// 以下の特徴を持つloggerを返す。
//  1. 出力先がファイル（ファイル名はfilePath）
//  2. ローテーションあり（100MBのログファイルごと）
//  3. 保存期間、保存ログ数は無制限
//
// 呼び出しもとでは使用後にwをClose()する必要がある。
func NewLogger(filePath string) (l *slog.Logger, w *lumberjack.Logger) {
	w = &lumberjack.Logger{
		Filename:  filePath,
		LocalTime: true,
	}
	l = slog.New(slog.NewJSONHandler(w, nil))
	return
}
