package main

import (
	"github.com/Jacobbrewer1/bindicator/bins"
	"github.com/Jacobbrewer1/bindicator/config"
	"github.com/Jacobbrewer1/bindicator/email"
	"log"
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
			b, err := bins.GetBins(*p.UPRN)
			if err != nil {
				log.Println(err)
				continue
			}
			name, s := b.NextBin()
			email.WaitAndSend(name, s, p)
		}
		time.Sleep(time.Hour * 72)
	}
}

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
	}
	go run()
	time.Sleep(time.Minute)
}
