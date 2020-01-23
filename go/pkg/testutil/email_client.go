package testutil

import (
	"fmt"
	"net/mail"

	"github.com/inbucket/inbucket/pkg/rest/client"

	"github.com/mmrath/gobase/go/pkg/email"
	"github.com/mmrath/gobase/go/pkg/errutil"
)

type TestEmailClient struct {
	client *client.Client
}

func NewEmailClient(mailAPI string) *TestEmailClient {
	c, err := client.New(mailAPI)
	if err != nil {
		fmt.Println("failed to create test email client")
		panic(err)
	}
	return &TestEmailClient{client: c}
}

func (c *TestEmailClient) GetLatestEmail(emailID string) *email.Message {
	headers, err := c.client.ListMailbox(emailID)
	if err != nil {
		return nil
	}
	for _, h := range headers {
		msg, err := h.GetMessage()
		if err != nil {
			panic(errutil.Wrap(err, "failed to load test emails"))
		}
		if msg != nil {
			return &email.Message{
				To:      toAddress(h.To...),
				From:    email.Address{Address: h.From},
				Subject: h.Subject,
				HTML:    msg.Body.HTML,
				Text:    msg.Body.Text,
			}
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
