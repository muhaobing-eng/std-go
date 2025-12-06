package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

func NewMailClient(from string, password string, to []string, subject string) *MailClient {
	mail := email.NewEmail()
	mail.From = fmt.Sprintf("%s <%s>", from, from)
	mail.To = to
	mail.Subject = subject
	return &MailClient{from: from, password: password, mail: mail}
}

type MailClient struct {
	from     string
	password string
	mail     *email.Email
}

func (m *MailClient) SetAlias(alias string) *MailClient {
	m.mail.From = fmt.Sprintf("%s <%s>", alias, m.from)
	return m
}

func (m *MailClient) SetCc(cc []string) *MailClient {
	m.mail.Cc = cc
	return m
}

func (m *MailClient) SetBcc(bcc []string) *MailClient {
	m.mail.Bcc = bcc
	return m
}

func (m *MailClient) SetContent(content string, isHtml bool) *MailClient {
	if isHtml {
		m.mail.HTML = []byte(content)
	} else {
		m.mail.Text = []byte(content)
	}
	return m
}

func (m *MailClient) SetAttachments(attachments []string) *MailClient {
	for _, file := range attachments {
		m.mail.AttachFile(file)
	}
	return m
}

func (m *MailClient) SendMail() error {
	suffix := strings.Split(m.from, "@")[1]
	server := SupportLists[suffix]
	auth := smtp.PlainAuth("", m.from, m.password, server.Host)
	return m.mail.Send(fmt.Sprintf("%s:%d", server.Host, server.Port), auth)
}
