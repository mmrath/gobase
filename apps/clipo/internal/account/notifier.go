package account

import (
	"fmt"
	"html/template"

	"github.com/mmrath/gobase/apps/clipo/internal/templateutil"
	"github.com/mmrath/gobase/pkg/errutil"

	"github.com/mmrath/gobase/pkg/email"
	"github.com/mmrath/gobase/pkg/model"
)

type Notifier interface {
	NotifyActivation(user model.User, token string) error
	NotifyPasswordChange(user model.User) error
	NotifyPasswordResetInit(user model.User, token string) error
}

func NewNotifier(appDomainName string, mailer email.Mailer, registry *templateutil.Registry) Notifier {
	return &notifier{appDomainName: appDomainName, mailer: mailer, templateRegistry: registry}
}

type notifier struct {
	mailer           email.Mailer
	appDomainName    string
	templateRegistry *templateutil.Registry
}

func (n *notifier) NotifyPasswordChange(user model.User) error {

	data := make(map[string]interface{})
	data["user"] = user

	from := email.NewAddress("", "")
	to := []email.Address{email.NewAddress(user.GetName(), user.GetEmail())}
	subject := "Account password changed"
	pcTmpl := "templates/email/auth/password_changed.gohtml"

	htmlBody, err := n.templateRegistry.RenderToString(pcTmpl, data)

	if err != nil {
		return errutil.Wrapf(err, "failed to render email")
	}

	msg, err := email.NewHTMLMessage(from, to, subject, htmlBody)
	if err != nil {
		return errutil.Wrap(err, "failed to build email message")
	}
	err = n.mailer.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (n *notifier) NotifyActivation(user model.User, token string) error {

	url := fmt.Sprintf("%s/clipo/account/activate?key=%s", n.appDomainName, token)
	data := struct {
		URL  template.URL
		User model.User
	}{
		URL:  template.URL(url),
		User: user,
	}

	from := email.NewAddress("info", "info@"+n.appDomainName)
	to := []email.Address{email.NewAddress(user.GetName(), user.GetEmail())}
	subject := "Activate your account"

	htmlBody, err := n.templateRegistry.RenderToString("templates/email/auth/account_activation.gohtml", data)

	if err != nil {
		return errutil.Wrapf(err, "failed to render email")
	}

	msg, err := email.NewHTMLMessage(from, to, subject, htmlBody)

	if err != nil {
		return errutil.Wrapf(err, "failed to create email message")
	}

	err = n.mailer.Send(msg)
	if err != nil {
		return errutil.Wrapf(err, "failed to send email")
	}
	return nil
}

func (n *notifier) NotifyPasswordResetInit(user model.User, token string) error {
	url := fmt.Sprintf("%s/account/reset-password?key=%s", n.appDomainName, token)
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

	htmlBody, err := n.templateRegistry.RenderToString(passwordResetTmpl, data)

	if err != nil {
		return errutil.Wrapf(err, "failed to render email")
	}

	msg, err := email.NewHTMLMessage(from, to, subject, htmlBody)

	if err != nil {
		return err
	}

	err = n.mailer.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
