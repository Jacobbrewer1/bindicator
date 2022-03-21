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
						name, s := person.NextBin()
						go email.WaitAndSend(name, s, person)
					}(p)
				} else {
					log.Printf("%v does not have any bins tomorrow\n", *p.Name)
				}
			}
		}()
		t := time.Now().UTC().Add(time.Hour * 24).Format(helper.TimeLayout)
		elms := strings.Split(t, "T")
		w, err := time.Parse(helper.TimeLayout, elms[0]+"T05:00:00Z")
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
