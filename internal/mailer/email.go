package mailer

import (
	"bytes"
	"fmt"

	"net/smtp"

	"github.com/spf13/viper"
	"github.com/the-Jinxist/tukio-api/config"
)

func SendEmail(toEmail, subject string, templateName string, templateVar interface{}) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	var body bytes.Buffer

	fromEmail := viper.GetString(config.FromEmailKey)
	fromPassword := viper.GetString(config.FromPasswordKey)
	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

	// tempPath := fmt.Sprintf("./templates/%s.html", templateName)
	// t, err := template.ParseFiles(tempPath)

	// if err != nil {
	// 	return err
	// }

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))

	// if err = t.Execute(&body, templateVar); err != nil {
	// 	return err
	// }

	addr := smtpHost + ":" + smtpPort

	err := smtp.SendMail(addr, auth, fromEmail, []string{toEmail}, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
