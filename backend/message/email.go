package message

import (
	"fmt"
	"nebula/config"
	"net/smtp"
)

// SendEmail .
func SendEmail(msg []byte, sendToAddress string) {
	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", config.EmailAccount,
		config.EmailPassword, hostname)

	err := smtp.SendMail(hostname+":587", auth, config.EmailAccount, []string{sendToAddress}, msg)
	if err != nil {
		fmt.Println("Failed to send email")
	}
}
