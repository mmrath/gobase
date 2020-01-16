// Package email provides email sending functionality.
package email

import (
	"fmt"
	"github.com/go-mail/mail"
	"github.com/hashicorp/errwrap"
	"github.com/jaytaylor/html2text"
	"github.com/vanng822/go-premailer/premailer"
	"html/template"
	"log"
	stdMail "net/mail"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	debug     bool
	templates *template.Template
)

type Mailer interface {
	Send(email *Message) error
}

// Mailer is a SMTP mailer.
type mailer struct {
	client *mail.Dialer
	from   Address
}

// NewMailer returns a configured SMTP Mailer.
func NewMailer(conf SMTPConfig) (Mailer, error) {

	s := &mailer{
		client: mail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password),
		from:   conf.From,
	}

	if conf.Host == "" {
		log.Println("SMTP host not set => printing emails To stdout")
		debug = true
		return s, nil
	}

	d, err := s.client.Dial()
	if err == nil {
		d.Close()
		log.Println("connected to mail server")
		return s, nil
	}

	return nil, errwrap.Wrapf("failed to dial mail server", err)
}

// Send sends the mail via smtp.
func (m *mailer) Send(email *Message) error {
	if debug {
		log.Println("To:", email.To)
		log.Println("Subject:", email.Subject)
		log.Println(email.Text)
		return nil
	}

	msg := mail.NewMessage()

	addresses := make([]string, len(email.To))
	for i := range addresses {
		addresses[i] = msg.FormatAddress(email.To[i].Address, email.To[i].Name)
	}

	msg.SetAddressHeader("From", email.From.Address, email.From.Name)
	msg.SetHeader("To", addresses...)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("Text/plain", email.Text)
	msg.AddAlternative("Text/Html", email.Html)

	return m.client.DialAndSend(msg)
}

// message struct holds all parts of a specific email message.
type Message struct {
	From     Address
	To       []Address
	Subject  string
	Html     string
	Text     string
}

func NewHtmlMessage(from Address, to []Address, subject string, htmlBody string) (*Message, error) {
	msg := Message{From: from, To: to, Subject: subject, Html: htmlBody}
	if err := msg.parse(); err != nil {
		return nil, err
	}
	return &msg, nil
}

// parse parses the corresponding Template and content
func (m *Message) parse() error {

	prem, err := premailer.NewPremailerFromString(m.Html, premailer.NewOptions())
	if err != nil {
		return err
	}
	html, err := prem.Transform()
	if err != nil {
		return err
	}
	m.Html = html

	text, err := html2text.FromString(html, html2text.Options{PrettyTables: true})
	if err != nil {
		return err
	}
	m.Text = text
	return nil
}

// Address struct holds email address and recipient name.
type Address = stdMail.Address

// NewEmail returns an email address.
func NewAddress(name string, address string) Address {
	return Address{
		Name:    name,
		Address: address,
	}
}

func LoadTemplates(templatePath string) error {
	if len(templatePath) == 0 {
		templatePath = "./resources/templates/email"
	}
	templates = template.New("").Funcs(fMap)
	return filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templates.ParseFiles(path)
			return err
		}
		return err
	})
}

var fMap = template.FuncMap{
	"formatAsDate":     formatAsDate,
	"formatAsDuration": formatAsDuration,
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d.%d.%d", day, month, year)
}

func formatAsDuration(t time.Time) string {
	dur := t.Sub(time.Now())
	hours := int(dur.Hours())
	mins := int(dur.Minutes())

	v := ""
	if hours != 0 {
		v += strconv.Itoa(hours) + " hours and "
	}
	v += strconv.Itoa(mins) + " minutes"
	return v
}
