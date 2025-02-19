package infrastructure

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)


const (
    SmtpHost      = "smtp.gmail.com"        
    SmtpPort      = 465                     
    EmailFrom     = "bezueyerusalem@gmail.com" 
    EmailPassword = "fxtf szty wpmy sznn"        
    ServerHost    = "http://localhost:8080"   
    TokenTTlL      = time.Hour               
)

// Generates a secure random token
func GenerateVerificationToken() (string, error) {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// Sends the password reset email
func SendVerificationEmail(to, token string) error {
	verificationLink := fmt.Sprintf("%s/api/auth/verify-email?token=%s&email=%s", ServerHost, token, to)
    body := fmt.Sprintf(`
        Hi,
		Please Verify Your Account by clicking the link below:
        %s

        If you did not request this, please ignore this email.
    `, verificationLink)

    m := gomail.NewMessage()
    m.SetHeader("From", fmt.Sprintf("%s <%s>", "Unique Minds", EmailFrom))
    m.SetHeader("To", to)
    m.SetHeader("Subject", "Email Verification")
    m.SetBody("text/plain", body)

    d := gomail.NewDialer(SmtpHost, SmtpPort, EmailFrom, EmailPassword)
    d.SSL = true

    if err := d.DialAndSend(m); err != nil {
        return err
    }


    return nil
}


// Sends the password reset email
func SendResetPasswordVerificationEmail(to, token string) error {
	verificationLink := fmt.Sprintf("%s/users/password-update?token=%s&email=%s", ServerHost, token, to)
    body := fmt.Sprintf(`
        Hi,
		Please Click the link below to reset your password:
        %s

        If you did not request this, please ignore this email.
    `, verificationLink)

    m := gomail.NewMessage()
    m.SetHeader("From", fmt.Sprintf("%s <%s>", "Eyerusalem Loan Tracking Project", EmailFrom))
    m.SetHeader("To", to)
    m.SetHeader("Subject", "Reset Password Verification")
    m.SetBody("text/plain", body)

    d := gomail.NewDialer(SmtpHost, SmtpPort, EmailFrom, EmailPassword)
    d.SSL = true

    if err := d.DialAndSend(m); err != nil {
        return err
    }
    return nil
}