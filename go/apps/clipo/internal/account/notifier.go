package account

import (
	"fmt"
	"github.com/mmrath/gobase/go/apps/clipo/internal/template_util"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"html/template"

	"github.com/mmrath/gobase/go/pkg/email"
	"github.com/mmrath/gobase/go/pkg/model"
)

type Notifier interface {
	NotifyActivation(user model.User, token string) error
	NotifyPasswordChange(user model.User) error
	NotifyPasswordResetInit(user model.User, token string) error
}

func NewNotifier(baseUrl string, mailer email.Mailer, registry *template_util.Registry) Notifier {
	return &notifier{baseUrl: baseUrl, sender: mailer, templateRegistry: registry}
}

type notifier struct {
	sender           email.Mailer
	baseUrl          string
	templateRegistry *template_util.Registry
}

func (e *notifier) NotifyPasswordChange(user model.User) error {

	data := make(map[string]interface{})
	data["user"] = user

	from := email.NewAddress("", "")
	to := []email.Address{email.NewAddress(user.GetName(), user.GetEmail())}
	subject := "Account password changed"
	pcTmpl := "templates/email/auth/password_changed.gohtml"

	htmlBody, err := e.templateRegistry.RenderToString(pcTmpl, data)

	if err != nil {
		return errutil.Wrapf(err, "failed to render email")
	}

	msg, err := email.NewHtmlMessage(from, to, subject, htmlBody)

	err = e.sender.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (e *notifier) NotifyActivation(user model.User, token string) error {

	url := fmt.Sprintf("%s/account/activate?key=%s", e.baseUrl, token)
	data := struct {
		URL  template.URL
		User model.User
	}{
		URL:  template.URL(url),
		User: user,
	}

	from := email.NewAddress("", "")
	to := []email.Address{email.NewAddress(user.GetName(), user.GetEmail())}
	subject := "Activate your account"

	htmlBody, err := e.templateRegistry.RenderToString("templates/email/auth/account_activation.gohtml", data)

	if err != nil {
		return errutil.Wrapf(err, "failed to render email")
	}

	msg, err := email.NewHtmlMessage(from, to, subject, htmlBody)

	if err != nil {
		return errutil.Wrapf(err, "failed to create email message")
	}

	err = e.sender.Send(msg)
	if err != nil {
		return errutil.Wrapf(err, "failed to send email")
	}
	return nil
}

func (e *notifier) NotifyPasswordResetInit(user model.User, token string) error {
	url := fmt.Sprintf("%s/account/reset-password?key=%s", e.baseUrl, token)
	data := struct {
		URL  template.URL
		User model.User
	}{
		URL:  template.URL(url),
		User: user,
	}

	from := email.NewAddress("test", "test@localhost")
	to := []email.Address{email.NewAddress(user.GetName(), user.GetEmail())}
	subject := "Reset password"
	passwordResetTmpl := "templates/email/auth/init_password_reset.gohtml"

	htmlBody, err := e.templateRegistry.RenderToString(passwordResetTmpl, data)

	if err != nil {
		return errutil.Wrapf(err, "failed to render email")
	}

	msg, err := email.NewHtmlMessage(from, to, subject, htmlBody)

	if err != nil {
		return err
	}

	err = e.sender.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
