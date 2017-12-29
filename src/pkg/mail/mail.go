package mail

import "net/smtp"

type Mail struct {
	From     string
	Username string
	Password string
	To       string
	Sub      string
	Msg      string
}

func (m *Mail) body() string {
	return "To: " + m.To + "\r\n" +
		"Subject: " + m.Sub + "\r\n\r\n" +
		m.Msg + "\r\n"
}

func GmailSend(m *Mail) error {
	smtpSvr := "smtp.gmail.com:587"
	auth := smtp.PlainAuth(
		"",
		m.Username,
		m.Password,
		"smtp.gmail.com",
	)

	err := smtp.SendMail(smtpSvr, auth, m.From, []string{m.To}, []byte(m.body()))
	if err != nil {
		return err
	}

	return nil
}
