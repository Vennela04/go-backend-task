package logger

import (
	"go.uber.org/zap"
)

// NewLogger initializes and returns a new Zap logger
func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
