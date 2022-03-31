package main

import (
	"github.com/Jacobbrewer1/bindicator/bins"
	"github.com/Jacobbrewer1/bindicator/config"
	"github.com/Jacobbrewer1/bindicator/email"
	"github.com/Jacobbrewer1/bindicator/helper"
	"log"
	"strings"
	"time"
)

func init() {
	log.Println("initializing logging")
	//log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("Logging initialized")
}

func run() {
	for {
		go func() {
			for _, p := range config.JsonConfigVar.RemoteConfig.People {
				bins.GetBins(p)
				if p.BinTomorrow() {
					go func(person *config.PeopleConfig) {
						log.Printf("%v has a bin tomorrow\n", *person.Name)
						s := person.GetBinsTomorrow()
						go email.WaitAndSend(s, person)
					}(p)
				} else {
					log.Printf("%v does not have any bins tomorrow\n", *p.Name)
				}
			}
		}()
		t := time.Now().UTC().Add(time.Hour * 24).Format(helper.TimeLayout)
		date := strings.Split(t, "T")[0]
		w, err := time.Parse(helper.TimeLayout, date+"T05:00:00Z")
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("waiting until %v to run again\n", w)
		time.Sleep(helper.CalculateTimeDifference(w))
		config.UpdateConfig()
	}
}

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
	}
	go run()
	<-make(chan struct{})
}
