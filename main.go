package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	// Add this line to import the util package
	"github.com/joho/godotenv"
	// Replace "your-package-path" with the actual package path
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	from := os.Getenv("email")
	pass := os.Getenv("PASS_EMAIL")
	// log.Println(from)
	// log.Println("PASS_EMAIL:", os.Getenv("PASS_EMAIL"))
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	to := []string{"tind.nd6999@gmail.com"}

	us_password := util.GeneratePass(10)

	message := fmt.Sprintf("Subject: primetokenlist registration\n\nThanks for choosing us platform. Your password is: %s\nUse his for auth on our site with your email", us_password)
	log.Println(string([]byte(message)))

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")

}
