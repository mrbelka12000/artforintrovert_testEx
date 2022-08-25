// Package logger implements our custom logger.
package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Interface interface {
	Debug(msg string)
	Debugf(msg string, args ...interface{})
	Error(msg string)
	Errorf(msg string, args ...interface{})
	Info(msg string)
	Infof(msg string, args ...interface{})
	Warn(msg string)
	Warnf(msg string, args ...interface{})
}

type Logger struct {
	logger *zap.SugaredLogger
}

var _ Interface = (*Logger)(nil)

func NewLogger() (*Logger, error) {
	cfg := zap.NewDevelopmentConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("could not build logger: %w", err)
	}

	return &Logger{
		logger: logger.Sugar(),
	}, nil
}

//Sync clears log buffer
func (l *Logger) Sync() error {
	return l.logger.Sync()
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args)
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.logger.Infof(msg, args)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args)
}
