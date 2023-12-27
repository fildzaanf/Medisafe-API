package mailer

import (
	"strconv"
	"strings"
	"talkspace/app/configs"

	"github.com/go-mail/mail"
	"github.com/sirupsen/logrus"
)

func SendEmail(to []string, subject, template string, data interface{}) (bool, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load smtp configuration: %v", err)
	}

	m := mail.NewMessage()
	m.SetHeader("From", config.SMTP.SMTP_USER)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)

	emailContent := strings.Replace(template, "{{.notification}}", data.(string), -1)

	m.SetBody("text/html", emailContent)

	SMTP_PORT, err := strconv.Atoi(config.SMTP.SMTP_PORT)
	if err != nil {
		return false, err
	}

	d := mail.NewDialer(
		config.SMTP.SMTP_HOST,
		SMTP_PORT,
		config.SMTP.SMTP_USER,
		config.SMTP.SMTP_PASS,
	)

	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}

func SendOTPViaEmail(to []string, subject, template string, data interface{}) (bool, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load smtp configuration: %v", err)
	}

	m := mail.NewMessage()
	m.SetHeader("From", config.SMTP.SMTP_USER)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)

	emailContent := strings.Replace(template, "{{.otp}}", data.(string), -1)

	m.SetBody("text/html", emailContent)

	SMTP_PORT, err := strconv.Atoi(config.SMTP.SMTP_PORT)
	if err != nil {
		return false, err
	}

	d := mail.NewDialer(
		config.SMTP.SMTP_HOST,
		SMTP_PORT,
		config.SMTP.SMTP_USER,
		config.SMTP.SMTP_PASS,
	)

	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}
