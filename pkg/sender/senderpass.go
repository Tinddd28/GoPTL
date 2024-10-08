package sender

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	// Add this line to import the util package
	"github.com/joho/godotenv"
)

type Sender struct {
	Email    string
	Pass     string
	Name     string
	LastName string
}

func SendMail(s Sender) error {
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
	to := []string{s.Email}

	message := fmt.Sprintf("Subject: PrimeTokenList Registration\n\nDear %s %s!\nThanks for choosing us platform. Your password is: %s\nUse him for auth on our site with your email",
		s.LastName, s.Name, s.Pass)

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}

func SendMesResPass(password, email string) error {
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
	to := []string{email}

	message := fmt.Sprintf("Subject: PrimeTokenList Reset Password\n\nYour new password is: %s\nUse him for auth on our site with your email",
		password)

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}

func SendVerification(id int, email string) error {
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
	to := []string{email}

	//code := random.RandomInt(10000, 99999)
	message := fmt.Sprintf("Subject: PrimeTokenList Verifcation\n\nClick on the link to verify your email: http://0.0.0.0:8000/user/verification_accept/%d", id)

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
