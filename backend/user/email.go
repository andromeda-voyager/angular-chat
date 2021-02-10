package user

import (
	"fmt"
	"nebula/config"
	"net/smtp"
)

// SendEmail .
func sendEmail(msg []byte, sendToAddress string) {
	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", config.EmailAccount,
		config.EmailPassword, hostname)

	err := smtp.SendMail(hostname+":587", auth, config.EmailAccount, []string{sendToAddress}, msg)
	if err != nil {
		fmt.Println("Failed to send email")
	}
}

// SendCodeToEmail .
func SendCodeToEmail(email string) {
	if IsEmailInUse(email) {
		sendEmail([]byte("An account already exists with this email."), email)
	} else {
		msg := []byte("Nebula\n\nVerifcation Code:\t" + generateCode(email))
		sendEmail(msg, email)
	}
}
