package main

import (
	"github.com/Jacobbrewer1/bindicator/bins"
	"github.com/Jacobbrewer1/bindicator/config"
	"github.com/Jacobbrewer1/bindicator/email"
	"log"
	"sync"
	"time"
)

func init() {
	log.Println("Initializing logging")
	//log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("Logging initialized")
}

func run() {
	for {
		for _, p := range config.JsonConfigVar.RemoteConfig.People {
			go func(person *config.PeopleConfig) {
				for {
					b, err := bins.GetBins(*person.UPRN)
					if err != nil {
						log.Println(err)
						return
					}
					name, s := b.NextBin()
					email.WaitAndSend(name, s, person)
				}
			}(p)
		}
		time.Sleep(time.Hour * 72)
	}
}

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
	}
	var w sync.WaitGroup
	w.Add(1)
	go run()
	w.Wait()
}
