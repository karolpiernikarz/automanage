package email

import (
	"time"

	"github.com/spf13/viper"
	mail "github.com/xhit/go-simple-mail/v2"
)

// Send sends email
func Send(to string, body string, subject string) (err error) {
	server := getConfig()
	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}
	email := mail.NewMSG()
	email.SetFrom(viper.GetString("app.company") + " <no-reply@" + viper.GetString("aws.smtpdomain") + ">").
		AddTo(to).
		SetSubject(subject)
	email.SetBody(mail.TextHTML, body)
	err = email.Send(smtpClient)
	return err
}

// SendWithAttachment sends email with attachment
func SendWithAttachment(to string, body string, subject string, filepath string, filename string, inline bool) (err error) {
	server := getConfig()
	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}
	email := mail.NewMSG()
	email.SetFrom(viper.GetString("app.company") + " <no-reply@" + viper.GetString("aws.smtpdomain") + ">").
		AddTo(to).
		SetSubject(subject)
	email.SetBody(mail.TextHTML, body)
	email.Attach(&mail.File{FilePath: filepath, Name: filename, Inline: inline})
	err = email.Send(smtpClient)
	return err
}

// SendWithAttachmentAndCC needs testing
func SendWithAttachmentAndCC(to string, cc string, body string, subject string, filepath string, filename string, inline bool) (err error) {
	server := getConfig()
	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}
	email := mail.NewMSG()
	email.SetFrom(viper.GetString("app.company") + " <no-reply@" + viper.GetString("aws.smtpdomain") + ">").
		AddTo(to).
		AddCc(cc).
		SetSubject(subject)
	email.SetBody(mail.TextHTML, body)
	email.Attach(&mail.File{FilePath: filepath, Name: filename, Inline: inline})
	err = email.Send(smtpClient)
	return err
}

// SendToMultiple sends email to multiple people
func SendToMultiple(to []string, body string, subject string) (err error) {
	server := getConfig()
	smtpClient, err := server.Connect()
	if err != nil {
		panic(err)
	}
	email := mail.NewMSG()
	email.SetFrom(viper.GetString("app.company") + " <no-reply@" + viper.GetString("aws.smtpdomain") + ">").
		AddTo(to...).
		SetSubject(subject)
	email.SetBody(mail.TextHTML, body)
	err = email.Send(smtpClient)
	return err
}

func getConfig() *mail.SMTPServer {
	server := mail.NewSMTPClient()

	server.Host = "email-smtp." + viper.GetString("aws.region") + ".amazonaws.com"
	server.Port = viper.GetInt("aws.mailport")
	server.Username = viper.GetString("aws.accesskeyid")
	server.Password = viper.GetString("aws.smtpsecretkey")
	server.Encryption = mail.EncryptionSTARTTLS // need to be get the value from config.yaml
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 60 * time.Second
	return server
}
