package utils

import (
	"access-management-system/config"
	"net/smtp"
	"strconv"

	"github.com/outbrain/golib/log"
)

// Send email verification code
func SendEmailVerificationCode(toMail string, verificationCode string) (err error) {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config", err)
	}

	from := config.EmailFrom
	appPassword := config.SMTPAppPassword

	to := []string{
		toMail,
	}

	smtpPort := strconv.Itoa(config.SMTPPort)

	addr := config.SMTPHost + ":" + smtpPort
	host := config.SMTPHost

	message := "Your email verification code is " + verificationCode

	msg := []byte("From: Access Management System\r\n" +
		"To: " + toMail + "\r\n" +
		"Subject: Your email verification code\r\n\r\n" +
		message + "\r\n")

	auth := smtp.PlainAuth("", from, appPassword, host)

	err = smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Error("Send mail error: ", err)
		return err
	}

	return nil
}

// Send admin approval email
func SendApprovedEmail(toMail string) (err error) {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config", err)
	}

	from := config.EmailFrom
	appPassword := config.SMTPAppPassword

	to := []string{
		toMail,
	}

	smtpPort := strconv.Itoa(config.SMTPPort)

	addr := config.SMTPHost + ":" + smtpPort
	host := config.SMTPHost

	message := "Welcome from access management system!!!!!"
	msg := []byte("From: Access Management System\r\n" +
		"To: " + toMail + "\r\n" +
		"Subject: Approved ams account\r\n\r\n" +
		message + "\r\n")
	auth := smtp.PlainAuth("", from, appPassword, host)

	err = smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Error("Send mail error: ", err)
		return err
	}

	return nil
}
