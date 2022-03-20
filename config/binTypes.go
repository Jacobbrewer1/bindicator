package config

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	RubbishText     = "Rubbish"
	FoodWasteText   = "Food Waste"
	RecyclingText   = "Recycling"
	GardenWasteText = "Garden Waste"
)

var (
	layout = "2006-01-02T15:04:05Z"
)

type Bins struct {
	Rubbish     *Rubbish     `json:"Rubbish"`
	FoodWaste   *FoodWaste   `json:"Food Waste"`
	Recycling   *Recycling   `json:"Recycling"`
	GardenWaste *GardenWaste `json:"Garden Waste"`
}

type Rubbish struct {
	*BinStruct
}

type FoodWaste struct {
	*BinStruct
}

type Recycling struct {
	*BinStruct
}

type GardenWaste struct {
	*BinStruct
}

type BinStruct struct {
	Previous *string `json:"Previous"`
	Next     *string `json:"Next"`
	PdfLink  *string `json:"PdfLink"`
	Communal *bool   `json:"Communal"`
}

func (b Bins) BinTomorrow() bool {
	_, x := b.NextBin()
	t := x.GetNextTime()
	y := time.Now().UTC().Format("2006-01-02")
	y = "2022-03-20T00:00:00Z"
	j, err := time.Parse(layout, y)
	if err != nil {
		log.Println(err)
		return false
	}
	j = j.Add(time.Hour * 24)
	return t.Equal(j)
}

func (b BinStruct) GetEmailTime() time.Time {
	rub, err := strconv.ParseInt(*b.Next, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	t := time.Unix(rub, 0).UTC()
	return t.Add(-time.Hour * 14)
}

func (b BinStruct) GetNextTime() time.Time {
	rub, err := strconv.ParseInt(*b.Next, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return time.Unix(rub, 0).UTC()
}

func (b BinStruct) GetNextTimeString() string {
	rub, err := strconv.ParseInt(*b.Next, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	yy, mm, dd := time.Unix(rub, 0).UTC().Date()
	return fmt.Sprintf("%v-%v-%v", dd, mm, yy)
}

func (b Bins) NextBin() (string, *BinStruct) {
	trub := b.Rubbish.GetNextTime()
	trec := b.Recycling.GetNextTime()
	tfw := b.FoodWaste.GetNextTime()
	tgw := b.GardenWaste.GetNextTime()

	if trub.Before(tgw) && trub.Before(trec) && trub.Before(tfw) {
		return RubbishText, b.Rubbish.BinStruct
	}
	if tgw.Before(trub) && tgw.Before(trec) && tgw.Before(tfw) {
		return GardenWasteText, b.GardenWaste.BinStruct
	}
	if trec.Before(trub) && trec.Before(tgw) && trec.Before(tfw) {
		return RecyclingText, b.Recycling.BinStruct
	}
	if tfw.Before(trub) && tfw.Before(tgw) && tfw.Before(trec) {
		return FoodWasteText, b.FoodWaste.BinStruct
	}
	return "", &BinStruct{Next: nil}
}

func (b *Bins) FormatBinDates() {
	var w sync.WaitGroup
	w.Add(4)
	go func(b *Bins) {
		defer w.Done()
		x := *b.Rubbish.Next
		*b.Rubbish.Next = x[strings.Index(x, "(")+1 : strings.Index(x, ")")-3]
		x = *b.Rubbish.Previous
		*b.Rubbish.Previous = x[strings.Index(x, "(")+1 : strings.Index(x, ")")-3]
	}(b)
	go func(b *Bins) {
		defer w.Done()
		x := *b.FoodWaste.Next
		*b.FoodWaste.Next = x[strings.Index(x, "(")+1 : strings.Index(x, ")")-3]
		x = *b.FoodWaste.Previous
		*b.FoodWaste.Previous = x[strings.Index(x, "(")+1 : strings.Index(x, ")")-3]
	}(b)
	go func(b *Bins) {
		defer w.Done()
		x := *b.Recycling.Next
		*b.Recycling.Next = x[strings.Index(x, "(")+1 : strings.Index(x, ")")-3]
		x = *b.Recycling.Previous
		*b.Recycling.Previous = x[strings.Index(x, "(")+1 : strings.Index(x, ")")-3]
	}(b)
	go func(b *Bins) {
		defer w.Done()
		x := *b.GardenWaste.Next
		*b.GardenWaste.Next = x[strings.Index(x, "(")+1 : strings.Index(x, ")")-3]
		x = *b.GardenWaste.Previous
		*b.GardenWaste.Previous = x[strings.Index(x, "(")+1 : strings.Index(x, ")")-3]
	}(b)
	w.Wait()
}
