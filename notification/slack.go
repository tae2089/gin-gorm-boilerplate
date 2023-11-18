package notification

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	// ErrorTitleText is error header of notification
	ErrorTitleText = "Server Error Occurred"
	ErrorColor     = "danger"
	UserName       = "Error Notification"
)

// Slack is a Slack notification service.
type Slack struct {
	url        string
	httpClient http.Client
}

// New returns instance of Slack.
func NewSlackProvider(httpClient http.Client, url string) *Slack {
	return &Slack{httpClient: httpClient, url: url}
}

// Message is representation of json message.
type Message struct {
	Username    string       `json:"username"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

// Attachment is representation of json attachment.
type Attachment struct {
	Fallback string  `json:"fallback"`
	Color    string  `json:"color"`
	Pretext  string  `json:"pretext"`
	Title    string  `json:"title"`
	Text     string  `json:"text"`
	Fields   []Field `json:"fields"`
	Footer   string  `json:"footer"`
}

// Field is representation of json field.
type Field struct {
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func (s *Slack) generateMessage(err error) Message {
	sm := Message{
		Username: UserName,
		Attachments: []Attachment{
			{
				Fallback: "",
				Color:    ErrorColor,
				Fields: []Field{
					{
						Value: err.Error(),
						Short: true,
					},
				},
			},
		},
	}

	return sm
}

func (s *Slack) Send(err error) error {
	slackMessage, err := json.Marshal(s.generateMessage(err))
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", s.url, bytes.NewBuffer(slackMessage))
	if err != nil {
		return err
	}

	resp, err := s.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()
	return nil
}
