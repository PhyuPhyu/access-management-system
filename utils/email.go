package utils

import (
	"access-management-system/config"
	"access-management-system/models"
	"net/smtp"
	"strconv"

	"github.com/outbrain/golib/log"
)

// Send email verification code
func SendEmailVerificationCode(user models.User, verificationCode string) (err error) {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config", err)
	}

	from := config.EmailFrom
	toMail := user.Email
	appPassword := config.SMTPAppPassword

	to := []string{
		toMail,
	}

	smtpPort := strconv.Itoa(config.SMTPPort)

	addr := config.SMTPHost + ":" + smtpPort
	host := config.SMTPHost

	message := "Hi " + user.Name + ",\n\n\nThanks for registering.\n\nPlease confirm your email address.\n\nVerification code is: " + verificationCode + "\n\n\nThanks,\nAMS"

	msg := []byte("From: Access Management System\r\n" +
		"To: " + toMail + "\r\n" +
		"Subject: Verify your email address\r\n\r\n" +
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
func SendApprovedEmail(user models.User) (err error) {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config", err)
	}

	from := config.EmailFrom
	toMail := user.Email
	appPassword := config.SMTPAppPassword

	to := []string{
		toMail,
	}

	smtpPort := strconv.Itoa(config.SMTPPort)

	addr := config.SMTPHost + ":" + smtpPort
	host := config.SMTPHost

	message := "Hi " + user.Name + ",\n\n\nThanks for registering.\n\nWelcome to access management system\n\n\nThanks,\nAMS"
	msg := []byte("From: Access Management System\r\n" +
		"To: " + toMail + "\r\n" +
		"Subject: Approved AMS account\r\n\r\n" +
		message + "\r\n")
	auth := smtp.PlainAuth("", from, appPassword, host)

	err = smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Error("Send mail error: ", err)
		return err
	}

	return nil
}
