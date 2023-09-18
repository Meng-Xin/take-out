package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type MySLog struct {
	slog *slog.Logger
}

func NewMySlog(setLevel string, filePath string) ILog {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile " + filePath)
		panic(err)
	}
	writer := io.MultiWriter(f, os.Stdout) // 文件 + 控制台输出
	level := new(slog.LevelVar)            // 设置日志等级
	switch setLevel {
	case "info":
		level.Set(slog.LevelInfo)
	case "debug":
		level.Set(slog.LevelDebug)
	case "warning":
		level.Set(slog.LevelWarn)
	case "error":
		level.Set(slog.LevelError)
	}
	handle := slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: level})
	log := slog.New(handle)
	return &MySLog{slog: log}
}

func (s *MySLog) Debug(args ...interface{}) {
	if len(args) == 1 {
		s.slog.Debug(args[0].(string))
	} else if len(args) > 1 {
		s.slog.Debug(args[0].(string), args[1:]...)
	}
}

func (s *MySLog) Info(args ...interface{}) {
	if len(args) == 1 {
		s.slog.Info(args[0].(string))
	} else if len(args) > 1 {
		s.slog.Info(args[0].(string), args[1:]...)
	}
}

func (s *MySLog) Warn(args ...interface{}) {
	if len(args) == 1 {
		s.slog.Warn(args[0].(string))
	} else if len(args) > 1 {
		s.slog.Warn(args[0].(string), args[1:]...)
	}
}

func (s *MySLog) Error(args ...interface{}) {
	if len(args) == 1 {
		s.slog.Error(args[0].(string))
	} else if len(args) > 1 {
		s.slog.Error(args[0].(string), args[1:]...)
	}
}

func (s *MySLog) Fatal(args ...interface{}) {
	if len(args) == 1 {
		s.slog.Info(args[0].(string))
	} else if len(args) > 1 {
		s.slog.Info(args[0].(string), args[1:]...)
	}
}
