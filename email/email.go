package email

import (
	"fmt"
	"github.com/Jacobbrewer1/bindicator/config"
	"github.com/Jacobbrewer1/bindicator/helper"
	"gopkg.in/gomail.v2"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func WaitAndSend(binName string, bin *config.BinStruct, p *config.PeopleConfig) {
	log.Printf("%v : waiting for %v bin\n", *p.Name, binName)
	log.Printf("%v : sending email at %v\n", *p.Name, bin.GetEmailTime().Format(time.RFC1123))
	time.Sleep(helper.CalculateTimeDifference(bin.GetEmailTime()))
	go sendEmail(p, bin, binName)
}

func sendEmail(p *config.PeopleConfig, bin *config.BinStruct, binName string) {
	b := bin.GetNextTime().UTC().Format(time.RFC1123)
	elms := strings.Split(b, " ")
	elms = elms[:len(elms) - 2]
	binDate := strings.Join(elms, " ")
	m := gomail.NewMessage()
	m.SetHeader("From", *config.JsonConfigVar.RemoteConfig.Email.From)
	m.SetHeader("To", *p.Email)
	m.SetHeader("Subject", fmt.Sprintf("%v bin!\n", binName))
	m.Embed(fmt.Sprintf(filepath.Join("assets", "images", "%v.jpeg"), strings.ToLower(strings.Join(strings.Split(binName, " "), ""))))
	m.SetBody("text/html", fmt.Sprintf(`%v bin is being collected tomorrow </br> %v </br><img src="cid:%v.jpeg" alt="bin image" />`,
		binName,
		binDate,
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
