package tests

import (
	"fmt"
	"github.com/inbucket/inbucket/pkg/rest/client"
	"github.com/mmrath/gobase/go/pkg/email"
	"net/mail"
)

type TestEmailClient struct {
	client *client.Client
}

func NewTestEmailClient(mailApi string) *TestEmailClient {
	c, err := client.New(mailApi)
	if err != nil {
		fmt.Println("failed to create test email client")
		panic(err)
	}
	return &TestEmailClient{client: c}
}

func (c *TestEmailClient) GetLatestEmail(emailId string) *email.Message {
	headers, err := c.client.ListMailbox(emailId)
	if err != nil {
		return nil
	}
	for _, h := range headers {
		msg, err := h.GetMessage()
		if err != nil {
			return nil
		}
		return &email.Message{
			To:      toAddress(h.To...),
			From:    email.Address{Address: h.From},
			Subject: h.Subject,
			Html:    msg.Body.HTML,
			Text:    msg.Body.Text,
		}
	}
	return nil
}

func toAddress(ids ...string) []email.Address {
	names := make([]email.Address, 0, len(ids))
	for _, name := range ids {
		address, err := mail.ParseAddress(name)
		if err != nil {
			panic(fmt.Sprintf("not able to parse address %v", err))
		}
		names = append(names, *address)
	}
	return names
}
