package utils

import (
	"gopkg.in/gomail.v2"
)

func EmailSetup() {
	m := gomail.NewMessage()

	m.SetHeader("From", "mailstrap.vincentprincewill44@gmail.com")
	m.SetHeader("To", "mailstrap.vincentprincewill44@gmail.com")
	m.SetAddressHeader("Cc", "", "password")

	m.SetHeader("Subject", "Dolang Email Test")

}
