package role

import (
	"bistleague-be/model/config"
	"bytes"
	"fmt"
	"mime"
	"net/smtp"
	"strings"
	"text/template"
	"github.com/gofiber/fiber/v2/log"
)

type Repository struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		cfg: cfg,
	}
}

func (r *Repository) SendEmailText(to []string, subject, body string) error {
	from := r.cfg.SMTP.User
	password := r.cfg.SMTP.Password

	smtpHost := r.cfg.SMTP.Host
	smtpPort := r.cfg.SMTP.Port

	// Create the email message with the subject.
	message := []byte("Subject: " + mime.QEncoding.Encode("utf-8", subject) + "\r\n" +
		"To: " + strings.Join(to, ", ") + "\r\n" +
		"From: " + from + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		log.Error(err)
	}

	return err
}

func (r *Repository) SendEmailHTML(to []string, subject, templateHTML string, data interface{}) error {
	from := r.cfg.SMTP.User
	password := r.cfg.SMTP.Password

	smtpHost := r.cfg.SMTP.Host
	smtpPort := r.cfg.SMTP.Port

	auth := smtp.PlainAuth("", from, password, smtpHost)
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))

	t, err := template.New("email").Parse(templateHTML)
	if err != nil {
		return err
	}

	err = t.Execute(&body, data)
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}
