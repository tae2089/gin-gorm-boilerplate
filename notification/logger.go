package notification

import "github.com/tae2089/bob-logging/logger"

type Logger struct{}

func NewLoggerProvider() SenderProvider {
	return &Logger{}
}

func (l *Logger) Send(err error) error {
	logger.Error(err)
	return nil
}
