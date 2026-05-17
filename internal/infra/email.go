package infra

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

type EmailClient struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailClient(host string, port int, user, pass string) *EmailClient {
	return &EmailClient{Host: host, Port: port, Username: user, Password: pass}
}

func (c *EmailClient) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", c.Username, c.Password, c.Host)
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", c.Username, to, subject, body))
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return smtp.SendMail(addr, auth, c.Username, []string{to}, msg)
}

func RenderTemplate(templateText string, data any) (string, error) {
	tmpl, err := template.New("email").Parse(templateText)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
