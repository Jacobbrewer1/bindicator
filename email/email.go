package email

import (
	"fmt"
	"github.com/Jacobbrewer1/bindicator/config"
	"gopkg.in/gomail.v2"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func WaitAndSend(binName string, bin *config.BinStruct, p *config.PeopleConfig) {
	log.Printf("%v : waiting for %v bin\n", *p.Name, binName)
	log.Printf("%v : sending email at %v\n", *p.Name, bin.GetEmailTime())
	//time.Sleep(calculateTimeDifference(bin.GetEmailTime()))
	sendEmail(p, bin, binName)
}

func sendEmail(p *config.PeopleConfig, bin *config.BinStruct, binName string) {
	// Authentication.
	/*auth := smtp.PlainAuth("", *config.JsonConfigVar.RemoteConfig.Email.From,
	*config.JsonConfigVar.RemoteConfig.Email.Password,
	*config.JsonConfigVar.RemoteConfig.Email.SmtpHost
	 */
	/*	// Sending email.
		err := smtp.SendMail(*config.JsonConfigVar.RemoteConfig.Email.SmtpHost+":"+*config.JsonConfigVar.RemoteConfig.Email.SmtpPort,
			auth, *config.JsonConfigVar.RemoteConfig.Email.From,
			to,
			[]byte(message))
		if err != nil {
			log.Println(err)
			return
		}*/

	m := gomail.NewMessage()
	m.SetHeader("From", *config.JsonConfigVar.RemoteConfig.Email.From)
	m.SetHeader("To", *p.Email)
	m.SetHeader("Subject", fmt.Sprintf("%v bin!\n", binName))
	m.Embed(fmt.Sprintf(filepath.Join("assets", "images", "%v.jpeg"), strings.ToLower(strings.Join(strings.Split(binName, " "), ""))))
	m.SetBody("text/html", fmt.Sprintf(`%v bin is being collected tomorrow </br> %v/%v/%v </br><img src="cid:%v.jpeg" alt="bin image" />`,
		binName,
		bin.GetNextTime().UTC().Day(),
		bin.GetNextTime().UTC().Month().String(),
		bin.GetNextTime().UTC().Year(),
		strings.ToLower(strings.Join(strings.Split(binName, " "), ""))))

	i, err := strconv.Atoi(*config.JsonConfigVar.RemoteConfig.Email.SmtpPort)
	if err != nil {
		log.Println(err)
		return
	}

	d := gomail.NewPlainDialer(*config.JsonConfigVar.RemoteConfig.Email.SmtpHost, i,
		*config.JsonConfigVar.RemoteConfig.Email.From, *config.JsonConfigVar.RemoteConfig.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v : email sent successfully", *p.Name)
}

func calculateTimeDifference(t time.Time) time.Duration {
	diff := t.Sub(time.Now())
	log.Println("time difference ", diff)
	return diff
}
