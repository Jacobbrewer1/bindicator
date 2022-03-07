package bins

import (
	"encoding/json"
	"github.com/Jacobbrewer1/bindicator/config"
	"io/ioutil"
	"net/http"
)

func GetBins(UPRN string) (*Bins, error) {
	rawJson, err := fetchBins(UPRN)
	if err != nil {
		return &Bins{}, err
	}
	b, err := decodeBins(rawJson)
	if err != nil {
		return &Bins{}, err
	}
	b.FormatBinDates()
	return &b, nil
}

func decodeBins(rawJson json.RawMessage) (Bins, error) {
	var b Bins
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
