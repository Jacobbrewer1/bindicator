package email

import (
	"fmt"
	"github.com/Jacobbrewer1/bindicator/config"
	"github.com/Jacobbrewer1/bindicator/helper"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
	"strings"
	"time"
)

func WaitAndSend(bin []*config.BinStruct, p *config.PeopleConfig) {
	log.Printf("%v : sending email at %v\n", *p.Name, p.GetEmailTime().Format(time.RFC1123))
	time.Sleep(helper.CalculateTimeDifference(p.GetEmailTime()))
	go sendEmail(p, bin)
}

func sendEmail(p *config.PeopleConfig, bin []*config.BinStruct) {
	log.Printf("%v : preparing email\n", p.Name)

	m := gomail.NewMessage()
	m.SetHeader("From", *config.JsonConfigVar.RemoteConfig.Email.From)
	m.SetHeader("To", *p.Email)
	m.SetHeader("Subject", "bin day!\n")

	// %v.jpeg
	// <img src="cid: {{ images }}" alt="bin image" />
	emailText := `The %v bin/s is being collected tomorrow </br> %v </br>`
	var binText string

	x := helper.GetTimeTomorrow().Format(time.RFC1123)
	elms := strings.Split(x, " ")
	elms = elms[:len(elms)-2]
	binDate := strings.Join(elms, " ")

	bodyText := "The date they are being collected on is " + binDate

	for _, b := range bin {
		binText = binText + *b.Name + ", "
	}
	binText = binText[:len(binText)-2]

	html := fmt.Sprintf(emailText,
		binText,
		bodyText)

	//m.Embed(fmt.Sprintf(filepath.Join("assets", "images", "%v.jpeg"), strings.ToLower(strings.Join(strings.Split(binName, " "), ""))))
	m.SetBody("text/html", html)

	i, err := strconv.Atoi(*config.JsonConfigVar.RemoteConfig.Email.SmtpPort)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%v : email prepared", *p.Name)
	log.Printf("%v : sending email", *p.Name)

	d := gomail.NewDialer(*config.JsonConfigVar.RemoteConfig.Email.SmtpHost, i,
		*config.JsonConfigVar.RemoteConfig.Email.From, *config.JsonConfigVar.RemoteConfig.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v : email sent successfully", *p.Name)
}
