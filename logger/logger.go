package main

import "go.uber.org/zap"

// должны сделать логгер который будет писать логи в файл
func NewLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}