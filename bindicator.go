package main

import (
	"github.com/Jacobbrewer1/bindicator/bins"
	"github.com/Jacobbrewer1/bindicator/config"
	"github.com/Jacobbrewer1/bindicator/email"
	"log"
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
						name, s := person.NextBin()
						email.WaitAndSend(name, s, person)
					}(p)
				}
			}
		}()
		time.Sleep(time.Hour * 24)
	}
}

func setup() {
	y := time.Now().UTC().Format("2006-01-02")
	y = y + "T00:00:00Z"
	j, err := time.Parse("2006-01-02T15:04:05Z", y)
	if err != nil {
		log.Println(err)
		return
	}
	j = j.Add(time.Hour * 24)
	diff := j.Sub(time.Now())
	log.Println("waiting to run at ", diff)
	time.Sleep(diff)
}

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
	}
	//setup() // TODO : Uncomment this out for PR
	go run()
	<-make(chan struct{})
}
