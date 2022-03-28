package config

import (
	"fmt"
	"github.com/Jacobbrewer1/bindicator/helper"
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
	Name     *string
	Previous *string `json:"Previous"`
	Next     *string `json:"Next"`
	PdfLink  *string `json:"PdfLink"`
	Communal *bool   `json:"Communal"`
}

func (b *Bins) SetupNames() {
	b.Rubbish.Name = helper.PointToString(RubbishText)
	b.Recycling.Name = helper.PointToString(RecyclingText)
	b.FoodWaste.Name = helper.PointToString(FoodWasteText)
	b.GardenWaste.Name = helper.PointToString(GardenWasteText)
}

func (b Bins) BinTomorrow() bool {
	x := b.GetBinsTomorrow()
	return len(x) > 0
}

func (b Bins) GetEmailTime() time.Time {
	t := helper.GetTimeTomorrow()
	return t.Add(-time.Hour * 14)
}

func (b BinStruct) GetNextTime() time.Time {
	rub, err := strconv.ParseInt(*b.Next, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return time.Unix(rub, 0).UTC()
}

func (b BinStruct) GetPreviousTime() time.Time {
	rub, err := strconv.ParseInt(*b.Previous, 10, 64)
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

func (b Bins) GetBinsTomorrow() []*BinStruct {
	trub := b.Rubbish.GetPreviousTime()
	trec := b.Recycling.GetPreviousTime()
	tfw := b.FoodWaste.GetPreviousTime()
	tgw := b.GardenWaste.GetPreviousTime()

	var binsReturn []*BinStruct
	t := helper.GetTimeTomorrow()
	if t.Equal(trub) {
		binsReturn = append(binsReturn, b.Rubbish.BinStruct)
	}
	if t.Equal(trec) {
		binsReturn = append(binsReturn, b.Recycling.BinStruct)
	}
	if t.Equal(tfw) {
		binsReturn = append(binsReturn, b.FoodWaste.BinStruct)
	}
	if t.Equal(tgw) {
		binsReturn = append(binsReturn, b.GardenWaste.BinStruct)
	}
	return binsReturn
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
