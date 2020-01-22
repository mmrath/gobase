package test_helper

import (
	"github.com/mmrath/gobase/go/pkg/email"
)

type MockMailer struct {
	messages []*email.Message
}

func NewMockMailer() (*MockMailer, error) {
	err := email.LoadTemplates("../resources/templates/email")
	if err != nil {
		return nil, err
	}
	return &MockMailer{messages: make([]*email.Message, 0, 10)}, nil
}

func (m *MockMailer) Send(msg *email.Message) error {
	m.messages = append(m.messages, msg)
	return nil
}

func (m *MockMailer) PopLastMessage() *email.Message {
	var msg *email.Message
	msg, m.messages = m.messages[len(m.messages)-1], m.messages[:len(m.messages)-1]
	return msg
}

type MailMessage struct {
	To      []string
	Subject []string
	Body    string
}
