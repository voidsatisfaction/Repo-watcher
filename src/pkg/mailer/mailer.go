package mailer

import (
	"Repo-watcher/src/pkg/config"
	"fmt"
	"net/smtp"
)

type Mailer struct {
	From       string
	Username   string
	Password   string
	To         []string
	smtpServer string
	auth       smtp.Auth
}

func New(c *config.ConfigFile) *Mailer {
	smtpServer := "smtp.gmail.com:587"
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("", c.Mail.Username, c.Mail.Password, host)
	return &Mailer{
		c.Mail.From,
		c.Mail.Username,
		c.Mail.Password,
		c.Mail.To,
		smtpServer,
		auth,
	}
}

func (m *Mailer) Send(body string) error {
	err := smtp.SendMail(m.smtpServer, m.auth, m.From, m.To, []byte(body))
	if err != nil {
		fmt.Printf("smtp error: %+v\n", err)
		return err
	}
	return nil
}
