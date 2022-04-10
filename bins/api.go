package bins

import (
	"encoding/json"
	"github.com/Jacobbrewer1/bindicator/config"
	"io/ioutil"
	"log"
	"net/http"
)

func GetBins(person *config.PeopleConfig) {
	rawJson, err := fetchBins(*person.UPRN)
	if err != nil {
		log.Println(err)
		return
	}
	b, err := decodeBins(rawJson)
	if err != nil {
		log.Println(err)
		return
	}
	b.FormatBinDates()
	person.Bins = b
	person.SetupNames()
}

func decodeBins(rawJson json.RawMessage) (config.Bins, error) {
	var b config.Bins
	err := json.Unmarshal(rawJson, &b)
	return b, err
}

func fetchBins(UPRN string) (json.RawMessage, error) {
	req, err := http.NewRequest(http.MethodGet, *config.JsonConfigVar.ConnectionStrings.BCPCouncil, nil)
	if err != nil {
		return nil, err
	}
	params := req.URL.Query()
	params.Add("UPRN", UPRN)
	req.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
