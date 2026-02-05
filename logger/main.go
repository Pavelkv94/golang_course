package main

import (
	"fmt"

	"go.uber.org/zap"
)

func foo(logger *zap.Logger) {
	fmt.Println("Hello, World!")
	logger.Error("Error message")

}
func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Hello, World!")

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warn message")

	foo(logger)
}



