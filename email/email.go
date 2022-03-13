package email

import (
	"fmt"
	"github.com/Jacobbrewer1/bindicator/config"
	"github.com/Jacobbrewer1/bindicator/templates"
	"log"
	"net/smtp"
	"time"
)

func WaitAndSend(binName string, bin *config.BinStruct, p *config.PeopleConfig) {
	log.Printf("%v : waiting for %v bin\n", *p.Name, binName)
	log.Printf("%v : sending email at %v\n", *p.Name, bin.GetEmailTime())
	//time.Sleep(calculateTimeDifference(bin.GetEmailTime()))
	sendEmail(p, createMessage(binName, *p, *bin))
}

func createMessage(binName string, person config.PeopleConfig, bin config.BinStruct) string {
	subject := fmt.Sprintf("Subject: %v bin!\n", binName)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := templates.Templates.DefinedTemplates()
	return fmt.Sprintf("%v%v%v", subject, mime, body)
}

func sendEmail(p *config.PeopleConfig, message string) {
	to := []string{*p.Email}
	// Authentication.
	auth := smtp.PlainAuth("", *config.JsonConfigVar.RemoteConfig.Email.From,
		*config.JsonConfigVar.RemoteConfig.Email.Password,
		*config.JsonConfigVar.RemoteConfig.Email.SmtpHost)

	// Sending email.
	err := smtp.SendMail(*config.JsonConfigVar.RemoteConfig.Email.SmtpHost+":"+*config.JsonConfigVar.RemoteConfig.Email.SmtpPort,
		auth, *config.JsonConfigVar.RemoteConfig.Email.From,
		to,
		[]byte(message))
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v : email sent successfully", *p.Name)
}

func calculateTimeDifference(t time.Time) time.Duration {
	diff := t.Sub(time.Now())
	log.Println("time difference ", diff)
	if diff < time.Hour*24 {
		return 0
	}
	return diff
}
