package account

import (
	"fmt"
	"html/template"

	"github.com/mmrath/gobase/pkg/email"
	"github.com/mmrath/gobase/model"
)

type Notifier interface {
	NotifyActivation(user *model.User, token string) error
	NotifyPasswordChange(user *model.User) error
	NotifyPasswordResetInit(user *model.User, token string) error
}

func NewNotifier(baseUrl string, mailer email.Mailer) Notifier {
	return &notifier{baseUrl: baseUrl, sender: mailer}
}

type notifier struct {
	sender  email.Mailer
	baseUrl string
}

func (e *notifier) NotifyPasswordChange(user *model.User) error {

	data := make(map[string]interface{})
	data["user"] = user

	from := email.NewAddress("", "")
	to := []email.Address{email.NewAddress(user.GetName(), user.GetEmail())}
	subject := "Account password changed"
	pcTmpl := "auth/password_changed.html"

	msg, err := email.NewMessage(from, to, subject, pcTmpl, &data)

	err = e.sender.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (e *notifier) NotifyActivation(user *model.User, token string) error {

	url := fmt.Sprintf("%s/account/activate?key=%s", e.baseUrl, token)
	data := struct {
		URL  template.URL
		User *model.User
	}{
		URL:  template.URL(url),
		User: user,
	}

	from := email.NewAddress("", "")
	to := []email.Address{email.NewAddress(user.GetName(), user.GetEmail())}
	subject := "Activate your account"
	acTmpl := "accountActivation"

	msg, err := email.NewMessage(from, to, subject, acTmpl, &data)

	if err != nil {
		return err
	}

	err = e.sender.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (e *notifier) NotifyPasswordResetInit(user *model.User, token string) error {
	url := fmt.Sprintf("%s/account/reset-password?key=%s", e.baseUrl, token)
	data := struct {
		URL  template.URL
		User *model.User
	}{
		URL:  template.URL(url),
		User: user,
	}

	from := email.NewAddress("", "")
	to := []email.Address{email.NewAddress(user.GetName(), user.GetEmail())}
	subject := "Reset password"
	passwordResetTmpl := "initPasswordReset"

	msg, err := email.NewMessage(from, to, subject, passwordResetTmpl, &data)

	if err != nil {
		return err
	}

	err = e.sender.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
