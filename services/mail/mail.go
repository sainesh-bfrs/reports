package mail

/*
 * File: mail.go
 * File Created: Tuesday, 12th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"net/smtp"
	"reports/config"
	"reports/helpers"
	"strings"
)

// Mailer ...
type Mailer struct {
	To      []string
	Cc      []string
	Subject string
	Body    string
}

// Send ...
func (m *Mailer) Send() {
	msg := "To: " + strings.Join(m.To, ",") + "\r\n"

	if len(m.Cc) > 0 {
		msg = msg + "cc: " + strings.Join(m.Cc, ",") + "\r\n"
	}

	msg = msg + "Subject: " + m.Subject + "\r\n"

	msg = msg + m.Body + "\r\n"

	host := config.Config.MailHost + ":" + config.Config.MailPort

	auth := smtp.PlainAuth("", config.Config.MailUsername, config.Config.MailPassword, config.Config.MailHost)

	err := smtp.SendMail(host, auth, config.Config.MailFrom, m.To, []byte(msg))

	helpers.LogError("Error in sending mail", err)
}
