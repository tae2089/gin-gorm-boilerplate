package notification

import (
	"net/http"
	"os"

	"github.com/tae2089/bob-logging/logger"
)

var errorChan chan error

func init() {
	errorChan = make(chan error)
}

type SenderProvider interface {
	Send(err error) error
}

type SenderProviderOption func() SenderProvider

func Listen(providers ...SenderProviderOption) {
	for e := range errorChan {
		for _, provider := range providers {
			err := provider().Send(e)
			if err != nil {
				logger.Error(err)
			}
		}
	}
}

func GetErrorChan() *chan error {
	return &errorChan
}

func SendError(err error) {
	errorChan <- err
}

func UseSlackProvider() SenderProvider {
	httpClient := http.Client{}
	slackWebookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if slackWebookURL == "" {
		panic("SLACK_WEBHOOK_URL is not set")
	}
	return NewSlackProvider(httpClient, slackWebookURL)
}

func UseLoggerProvider() SenderProvider {
	return &Logger{}
}
